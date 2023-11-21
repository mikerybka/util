package util

import "net/http"

func IsMutation(r *http.Request) bool {
	switch r.Method {
	case http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete:
		return true
	default:
		return false
	}
}
