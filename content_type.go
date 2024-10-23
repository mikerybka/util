package util

import (
	"net/http"
	"strings"
)

func ContentType(r *http.Request, format string) bool {
	for _, f := range strings.Split(r.Header.Get("Content-Type"), ",") {
		if f == format {
			return true
		}
	}
	return false
}
