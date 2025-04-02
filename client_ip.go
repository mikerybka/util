package util

import (
	"net"
	"net/http"
	"strings"
)

func ClientIP(r *http.Request) string {
	// Check X-Forwarded-For (can contain multiple IPs)
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		ips := strings.Split(fwd, ",")
		return strings.TrimSpace(ips[0]) // First IP is client
	}

	// Fallback to X-Real-IP
	if rip := r.Header.Get("X-Real-IP"); rip != "" {
		return rip
	}

	// Default to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
