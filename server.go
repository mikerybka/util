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
		Hosts:        make(map[string]http.Handler),
		TwilioClient: twilioClient,
		AdminPhone:   adminPhone,
	}
}

// Server hosts multiple apps.
// App data is read from "{datadir}/{host}/data.json".
type Server struct {
	DataDir      string
	Hosts        map[string]http.Handler
	TwilioClient *TwilioClient
	AdminPhone   string
}

func (s *Server) AddApp(host string, handler http.Handler) {
	s.Hosts[host] = handler
}

func (s *Server) HostPolicy(ctx context.Context, host string) error {
	host = strings.TrimPrefix(host, "www.")
	_, ok := s.Hosts[host]
	if ok {
		return nil
	}
	return fmt.Errorf("host not allowed: %s", host)
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

	// Find the handler for the host.
	h, ok := s.Hosts[strings.TrimPrefix(r.Host, "www.")]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Try to read from the data file
	datapath := filepath.Join(s.DataDir, r.Host, "data.json")
	ReadJSONFile(datapath, h)

	// Turn off CORS.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Server the request.
	h.ServeHTTP(w, r)

	// Write changes if necessary.
	if IsMutation(r) {
		err := WriteJSONFile(datapath, h)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Server) Start(email, certDir string) error {
	// Create a channel to receive errors from the HTTP servers.
	errChan := make(chan error)

	// Define the autocert manager.
	// See https://godoc.org/golang.org/x/crypto/acme/autocert#Manager for details.
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(certDir),
		HostPolicy: s.HostPolicy,
		Email:      email,
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
