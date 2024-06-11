package util

import "net/http"

type SinglePageApp struct {
	mainHTML []byte
	mainCSS  []byte
	mainJS   []byte
}

func (spa *SinglePageApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := ParsePath(r.URL.Path)
	if len(p) == 0 {

	}
}
