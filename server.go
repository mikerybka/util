package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"path/filepath"
	"runtime/debug"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

func NewServer(dataDir string, adminPhone string, twilioClient *TwilioClient) *Server {
	return &Server{
		DataDir:      dataDir,
		TwilioClient: twilioClient,
		AdminPhone:   adminPhone,
	}
}

// Server hosts multiple apps.
// App data is read from "{datadir}/{host}/data.json".
// Config is "{datadir}/{host}/config.json" defined by AppConfig.
type Server struct {
	DataDir      string
	TwilioClient *TwilioClient
	AdminPhone   string
	AdminEmail   string
	CertDir      string
}

func (s *Server) SystemdService() string {
	f := `[Unit]
	Description=Server
	After=network.target
	
	[Service]
	Type=simple
	Environment="TWILIO_ACCOUNT_SID=%s"
	Environment="TWILIO_AUTH_TOKEN=%s"
	Environment="TWILIO_PHONE_NUMBER=%s"
	Environment="DATA_DIR=%s"
	Environment="ADMIN_PHONE=%s"
	Environment="ADMIN_EMAIL=%s"
	Environment="CERT_DIR=%s"
	ExecStart=/usr/local/bin/server
	Restart=on-failure
	
	[Install]
	WantedBy=multi-user.target
`
	return fmt.Sprintf(f,
		s.TwilioClient.AccountSID,
		s.TwilioClient.AuthToken,
		s.TwilioClient.PhoneNumber,
		s.DataDir,
		s.AdminPhone,
		s.AdminEmail,
		s.CertDir,
	)
}

func (s *Server) HostPolicy(ctx context.Context, host string) error {
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle panics by notifying the admin.
	defer func() {
		if err := recover(); err != nil {
			// Log the error and stack trace
			log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
			// Respond with a 500 Internal Server Error
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()

	// Log the request.
	b, _ := httputil.DumpRequest(r, true)
	log.Println(string(b))

	// Handle www. redirects.
	if strings.HasPrefix(r.Host, "www.") {
		url := r.URL
		url.Host = strings.TrimPrefix("www.", url.Host)
		http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
		return
	}

	app := &App{
		Dir: filepath.Join(s.DataDir, r.Host),
	}
	app.ServeHTTP(w, r)
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