package util

import "net/http"

func ServeError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
