package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"

	_ "embed"

	"golang.org/x/crypto/acme/autocert"
)

// Server hosts multiple apps.
// App data is read from "{datadir}/{host}/data.json".
// Config is "{datadir}/{host}/config.json" defined by AppConfig.
type Server struct {
	DataFile          string
	ErrorsDir         string
	TwilioClient      *TwilioClient
	AdminPhone        string
	AdminEmail        string
	CertDir           string
	Users             *Table[*User]
	LoginCodes        map[string]string // user ID => 6 digit code
	SessionTokens     map[string]string // token => user ID
	LoginCodeMsgFmt   string
	GithubSourceDir   string
	AllowRegistration bool
}

func (s *Server) Load() {
	f, _ := os.Open(s.DataFile)
	json.NewDecoder(f).Decode(s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			ServeError(w, http.StatusInternalServerError)
			s.LogError(err)
			s.NotifyAdmin(fmt.Sprintf("ERROR: %s%s: %v", r.Host, r.URL.String(), err))
		}
	}()

	NewRequest(r).Log()

	switch r.URL.Path {
	case "/auth/register":
		s.RegisterForm().ServeHTTP(w, r)
	case "/auth/send-login-code":
		s.SendLoginCodeForm().ServeHTTP(w, r)
	case "/auth/login":
		s.LoginForm().ServeHTTP(w, r)
	case "/auth/logout":
		s.LogoutForm().ServeHTTP(w, r)
	case "/webhooks/github":
		s.GithubMirror().ServeHTTP(w, r)
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
		Name:        "server",
		Desc:        "server",
		After:       "network.target",
		Type:        "simple",
		Env:         []Pair[string, string]{},
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

func (s *Server) GithubMirror() *GithubMirror {
	return &GithubMirror{
		Workdir: s.GithubSourceDir,
	}
}

func (s *Server) RegisterForm() *Form {
	return &Form{
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
			id, err := s.Users.Insert(&User{
				Phone:     r.FormValue("Phone"),
				Email:     r.FormValue("Email"),
				FirstName: r.FormValue("First Name"),
				LastName:  r.FormValue("Last Name"),
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Redirect(w, r, "/auth/send-login-code?userID="+id, http.StatusFound)
		},
	}
}

func (a *Server) SendLoginCodeForm() *Form {
	return &Form{
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
			if a.LoginCodes == nil {
				a.LoginCodes = make(map[string]string)
			}
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
}

func (a *Server) LoginForm() *Form {
	return &Form{
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
			if a.SessionTokens == nil {
				a.SessionTokens = make(map[string]string)
			}
			a.SessionTokens[token] = userID

			http.SetCookie(w, &http.Cookie{
				Name:  "Token",
				Value: token,
				Path:  "/",
			})
			http.Redirect(w, r, "/", http.StatusFound)
		},
	}
}

func (a *Server) LogoutForm() *Form {
	return &Form{
		Name:   "Logout",
		Fields: []Field{},
		ServePOST: func(w http.ResponseWriter, r *http.Request) {
			token, err := r.Cookie("Token")
			if err == nil {
				delete(a.SessionTokens, token.Value)
				DeleteCookie(w, "Token")
			}
			http.Redirect(w, r, "/", http.StatusFound)
		},
	}
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

func (s *Server) LogError(e any) {
	stack := debug.Stack()
	err := os.MkdirAll(s.ErrorsDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filepath.Join(s.ErrorsDir, UnixNanoTimestamp()))
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(f, "%v\n%s\n", e, stack)
	if err != nil {
		panic(err)
	}
}

func (s *Server) NotifyAdmin(msg string) {
	err := s.TwilioClient.SendSMS(s.AdminPhone, msg)
	if err != nil {
		panic(err)
	}
}
