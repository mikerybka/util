package util

import "net/http"

func StartServer(s http.Handler) error {
	port := EnvVar("PORT", "3000")
	return http.ListenAndServe(":"+port, s)
}
