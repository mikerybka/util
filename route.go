package util

import (
	"net/http"
	"regexp"
)

// Route struct to store path and handler
type Route struct {
	Pattern *regexp.Regexp
	Handler http.HandlerFunc
}
