package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	_ "embed"

	"golang.org/x/crypto/acme/autocert"
)

func NewServer(dataFile string, adminPhone string, twilioClient *TwilioClient, certDir string) *Server {
	usersTable := NewTable[*User]()
	usersTable.AddUniqConstraint("ID")
	usersTable.AddUniqConstraint("Phone")
	usersTable.AddUniqConstraint("Email")
	return &Server{
		DataFile:        dataFile,
		TwilioClient:    twilioClient,
		AdminPhone:      adminPhone,
		Users:           usersTable,
		LoginCodes:      map[string]string{},
		SessionTokens:   map[string]string{},
		LoginCodeMsgFmt: "Your login code is %s",
		CertDir:         certDir,
	}
}

// Server hosts multiple apps.
// App data is read from "{datadir}/{host}/data.json".
// Config is "{datadir}/{host}/config.json" defined by AppConfig.
type Server struct {
	DataFile        string
	TwilioClient    *TwilioClient
	AdminPhone      string
	AdminEmail      string
	CertDir         string
	Users           *Table[*User]
	LoginCodes      map[string]string // user ID => 6 digit code
	SessionTokens   map[string]string // token => user ID
	LoginCodeMsgFmt string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle panics by notifying the admin.
	defer func() {
		if err := recover(); err != nil {
			// Log the error and stack trace
			log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
			// Respond with a 500 Internal Server Error
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			// Message the admin
			err := s.TwilioClient.SendSMS(s.AdminPhone, fmt.Sprintf("ERROR: %s: %v", r.URL.String(), err))
			if err != nil {
				panic(err)
			}
		}
	}()

	// Log the request.
	NewRequest(r).Log()

	// Handle auth pages
	switch r.URL.Path {
	case "/auth/register":
		s.Register(w, r)
	case "/auth/send-login-code":
		s.SendLoginCode(w, r)
	case "/auth/login":
		s.Login(w, r)
	// case "/auth/whoami":
	// 	s.WhoAmI(w, r)
	// 	return
	// case "/auth/logout":
	// 	s.Logout(w, r)
	// 	return
	default:
		s.User(r).ServeHTTP(w, r)
	}

	if IsMutation(r) {
		err := WriteJSONFile(s.DataFile, s)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Server) SystemdService() *SystemdService {
	return &SystemdService{
		Name:  "server",
		Desc:  "server",
		After: "network.target",
		Type:  "simple",
		Env: []Pair[string, string]{
			{
				K: "TWILIO_ACCOUNT_SID",
				V: s.TwilioClient.AccountSID,
			},
			{
				K: "TWILIO_AUTH_TOKEN",
				V: s.TwilioClient.AuthToken,
			},
			{
				K: "TWILIO_PHONE_NUMBER",
				V: s.TwilioClient.PhoneNumber,
			},
			{
				K: "DATA_FILE",
				V: s.DataFile,
			},
			{
				K: "ADMIN_PHONE",
				V: s.AdminPhone,
			},
			{
				K: "ADMIN_EMAIL",
				V: s.AdminEmail,
			},
			{
				K: "CERT_DIR",
				V: s.CertDir,
			},
		},
		AutoRestart: "on-failure",
		WantedBy:    "multi-user.target",
	}
}

func (s *Server) HostPolicy(ctx context.Context, host string) error {
	return nil
}

func (s *Server) User(r *http.Request) *User {
	token, err := r.Cookie("Token")
	if err != nil {
		return nil
	}
	userID, ok := s.SessionTokens[token.Value]
	if !ok {
		return nil
	}
	user, ok := s.Users.Find(userID)
	if !ok {
		return nil
	}
	return user
}

func (s *Server) Start() error {
	// Create a channel to receive errors from the HTTP servers.
	errChan := make(chan error)

	// Define the autocert manager.
	// See https://godoc.org/golang.org/x/crypto/acme/autocert#Manager for details.
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(s.CertDir),
		HostPolicy: s.HostPolicy,
		Email:      s.AdminEmail,
	}

	// Start the HTTP server.
	go func() {
		err := http.ListenAndServe(":80", m.HTTPHandler(s))
		if err != nil {
			errChan <- err
		}
	}()

	// Start the HTTPS server.
	go func() {
		err := http.Serve(m.Listener(), s)
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for an error.
	return <-errChan
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	form := &Form{
		Name: "Register",
		Fields: []Field{
			{
				Name: "First Name",
				Type: StringType,
			},
			{
				Name: "Last Name",
				Type: StringType,
			},
			{
				Name: "Phone",
				Type: StringType,
			},
			{
				Name: "Email",
				Type: StringType,
			},
		},
		ServePOST: func(w http.ResponseWriter, r *http.Request) {
			user := &User{
				ID:        RandomID(),
				Phone:     r.FormValue("Phone"),
				Email:     r.FormValue("Email"),
				FirstName: r.FormValue("First Name"),
				LastName:  r.FormValue("Last Name"),
			}
			err := s.Users.Insert(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			http.Redirect(w, r, "/auth/send-login-code", http.StatusFound)
		},
	}
	form.ServeHTTP(w, r)
}

// SendLoginCode sends a login code the users phone number on file via Twilio SMS.
func (a *Server) SendLoginCode(w http.ResponseWriter, r *http.Request) {
	form := &Form{
		Name: "Send Login Code",
		Fields: []Field{
			{
				Name: "Phone",
				Type: StringType,
			},
		},
		ServePOST: func(w http.ResponseWriter, r *http.Request) {
			phone := r.FormValue("Phone")

			// Find the user.
			users := a.Users.FindBy("Phone", phone)
			userID, user, found := OnlyOne(users)
			if !found {
				http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
				return
			}

			// Create a login code.
			code := RandomCode(6)
			a.LoginCodes[userID] = code

			// Send the login code.
			msg := fmt.Sprintf(a.LoginCodeMsgFmt, code)
			err := a.TwilioClient.SendSMS(user.Phone, msg)
			if err != nil {
				panic(err)
			}

			http.Redirect(w, r, "/auth/login?UserID="+userID, http.StatusFound)
		},
	}
	form.ServeHTTP(w, r)
}

// Login creates a user session.
// It returns a token or error.
func (a *Server) Login(w http.ResponseWriter, r *http.Request) {
	form := &Form{
		Name: "Login",
		Fields: []Field{
			{
				Name: "UserID",
				Type: StringType,
			},
			{
				Name: "Code",
				Type: StringType,
			},
		},
		ServePOST: func(w http.ResponseWriter, r *http.Request) {
			userID := r.FormValue("UserID")
			code := r.FormValue("Code")
			ok := a.CheckLoginCode(userID, code)
			if !ok {
				http.Error(w, "wrong code", http.StatusBadRequest)
				return
			}

			// Invalidate the login code
			delete(a.LoginCodes, userID)

			// Generate a unique token
			var token string
			for {
				token = RandomToken(32)
				if _, ok := a.SessionTokens[token]; !ok {
					break
				}
			}

			// Save the token.
			a.SessionTokens[token] = userID

			http.SetCookie(w, &http.Cookie{
				Name:  "Token",
				Value: token,
				Path:  "/",
			})
			http.Redirect(w, r, "/", http.StatusFound)
		},
	}
	form.ServeHTTP(w, r)
}

func (a *Server) WhoAmI(token string) string {
	return a.SessionTokens[token]
}

func (a *Server) Logout(token string) {
	delete(a.SessionTokens, token)
}

// Creates a new login code.
// Only one login code per user can be active at a time.
func (a *Server) CreateLoginCode(userID string) string {
	// Generate a 6 digit code.
	code := RandomCode(6)

	// Save the code.
	a.LoginCodes[userID] = code

	return code
}

// CheckLoginCode returns true if the login code is valid.
func (a *Server) CheckLoginCode(userID, code string) bool {
	if code == "" {
		return false
	}
	return a.LoginCodes[userID] == code
}
