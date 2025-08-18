package util

import (
	"fmt"
	"net/http"
)

func ErrorHTTP(w http.ResponseWriter, code int) {
	msg := ""
	switch code {
	case http.StatusMethodNotAllowed:
		msg = "method not allowed"
	default:
		panic(fmt.Sprintf("unknown code: %d", code))
	}
	http.Error(w, msg, code)
}
