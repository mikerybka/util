package util

import "net/http"

func Serve(s http.Handler) {
	panic(http.ListenAndServe(Port(), s))
}
