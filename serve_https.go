package util

import (
	"context"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

// Use Let's Encrypt to fetch and renew certificates on any domain.
// serveHTTPS binds to ports 80 and 443 and serves the given handler.
// It uses a special handler for port 80 that can handle ACME challenges.
func ServeHTTPS(s http.Handler, email, certDir string) error {
	// Create a channel to receive errors from the HTTP servers.
	errChan := make(chan error)

	// Define the autocert manager.
	// See https://godoc.org/golang.org/x/crypto/acme/autocert#Manager for details.
	m := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(certDir),
		HostPolicy: func(ctx context.Context, host string) error {
			return nil
		},
		Email: email,
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
