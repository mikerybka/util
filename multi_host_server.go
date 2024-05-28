package util

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

type MultiHostServer struct {
	Hosts map[string]http.Handler
}

func (s *MultiHostServer) HostPolicy(ctx context.Context, host string) error {
	_, ok := s.Hosts[host]
	if ok {
		return nil
	}
	return fmt.Errorf("host not allowed: %s", host)
}

func (s *MultiHostServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, ok := s.Hosts[r.Host]
	if !ok {
		http.NotFound(w, r)
		return
	}
	HandleCORS(w, r)
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
