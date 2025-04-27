package util

import "net/http"

func NewServeMux(handlers map[string]http.Handler) http.Handler {
	mux := http.NewServeMux()
	for k, v := range handlers {
		mux.Handle(k, v)
	}
	return mux
}
