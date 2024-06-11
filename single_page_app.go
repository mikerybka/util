package util

import "net/http"

type SinglePageApp struct {
	mainHTML []byte
	mainCSS  []byte
	mainJS   []byte
}

func (spa *SinglePageApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/main.css" {
		w.Write(spa.mainCSS)
		return
	}
	if r.URL.Path == "/main.js" {
		w.Write(spa.mainJS)
		return
	}
	w.Write(spa.mainHTML)
}
