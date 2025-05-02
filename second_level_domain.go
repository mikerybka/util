package util

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func SecondLevelDomain(r *http.Request) (string, error) {
	if strings.HasPrefix(r.Host, "localhost") {
		return "localhost", nil
	}
	// Remove port if present
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		// If error is missing port, keep the original host
		if !isMissingPort(err) {
			return "", err
		}
	}

	domain, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return "", fmt.Errorf("failed to extract second-level domain: %w", err)
	}
	return domain, nil
}

// Helper: check if error is just about missing port
func isMissingPort(err error) bool {
	_, ok := err.(*net.AddrError)
	return ok
}
