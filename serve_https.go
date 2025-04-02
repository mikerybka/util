package util

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

// ServeHTTPS serves h on HTTP and HTTPS ports handling tls.
func ServeHTTPS(h http.Handler, email, certDir string, allowHost func(host string) bool) error {
	// Create a channel to receive errors. Usually from ports already being used.
	errChan := make(chan error)

	// Define the autocert manager.
	// See https://godoc.org/golang.org/x/crypto/acme/autocert#Manager for details.
	m := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(certDir),
		HostPolicy: func(ctx context.Context, host string) error {
			if !allowHost(host) {
				return fmt.Errorf("host not allowed: %s", host)
			}
			return nil
		},
		Email: email,
	}

	// Start the HTTP server.
	go func() {
		err := http.ListenAndServe(":80", m.HTTPHandler(h))
		if err != nil {
			errChan <- err
		}
	}()

	// Start the HTTPS server.
	go func() {
		s := &http.Server{
			Addr:      ":443",
			Handler:   h,
			TLSConfig: m.TLSConfig(),
		}
		err := s.ListenAndServeTLS("", "")
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for an error.
	return <-errChan
}
