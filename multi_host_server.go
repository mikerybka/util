package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

type MultiHostServer struct {
	Hosts        map[string]http.Handler
	TwilioClient *TwilioClient
	AdminPhone   string
}

func (s *MultiHostServer) HostPolicy(ctx context.Context, host string) error {
	host = strings.TrimPrefix(host, "www.")
	_, ok := s.Hosts[host]
	if ok {
		return nil
	}
	return fmt.Errorf("host not allowed: %s", host)
}

func (s *MultiHostServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	// Turn off CORS.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Server the request.
	h.ServeHTTP(w, r)
}

func (s *MultiHostServer) Start(email, certDir string) error {
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
