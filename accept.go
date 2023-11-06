package util

import (
	"net/http"
	"strings"
)

func Accept(r *http.Request, format string) bool {
	for _, f := range strings.Split(r.Header.Get("Accept"), ",") {
		if f == format {
			return true
		}
	}
	return false
}
