package util

import "net/http"

type UI struct {
	Meta AppMeta
}

func (ui *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := &Router{
		Root: func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
			case "POST":
			}
		},
	}
	router.ServeHTTP(w, r)
}
