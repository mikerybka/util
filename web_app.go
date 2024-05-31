package util

import (
	"net/http"
	"strings"
)

type WebApp struct {
	Frontend *WebFrontend
	API      *WebAPI
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api") {
		http.StripPrefix("/api", app.API)
		return
	}
	app.Frontend.ServeHTTP(w, r)
}
