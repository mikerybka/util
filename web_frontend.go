package util

import "net/http"

var mainJS []byte
var mainHTML []byte

type WebFrontend struct {
	Favicon []byte
}

func (fe *WebFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.Write(fe.Favicon)
		return
	}
	if r.URL.Path == "/main.js" {
		w.Write(mainJS)
		return
	}
	w.Write(mainHTML)
}
