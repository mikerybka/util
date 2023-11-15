package util

import "net/http"

func Serve(s http.Handler) error {
	return http.ListenAndServe(GetAddr(), s)
}
