package util

import "net/http"

func WriteNotFound(w http.ResponseWriter) {
	http.Error(w, "not found", http.StatusNotFound)
}
