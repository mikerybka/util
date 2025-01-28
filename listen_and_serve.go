package util

import "net/http"

func ListenAndServe(h http.Handler) error {
	addr := ":" + RequireEnvVar("PORT")
	return http.ListenAndServe(addr, h)
}
