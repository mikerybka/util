package util

import "net/http"

func OptionallyServeHTTP(v any, w http.ResponseWriter, r *http.Request) {
	if handler, ok := v.(http.Handler); ok {
		handler.ServeHTTP(w, r)
	}
}
