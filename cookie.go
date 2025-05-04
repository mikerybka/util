package util

import (
	"net/http"
)

func Cookie(r *http.Request, id string) string {
	cookie, err := r.Cookie(id)
	if err != nil {
		return r.Header.Get(id)
	}
	return cookie.Value
}
