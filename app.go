package util

import (
	"net/http"
)

type App struct {
	DataDir  string
	Types    map[string]*Type
	RootType *Type
}

func (a App) HTMLReader() *HTMLReader {
	return NewHTMLReader(a.Types)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.Get(w, r)
	default:
		ErrorHTTP(w, http.StatusMethodNotAllowed)
	}
}

func (a *App) Get(w http.ResponseWriter, r *http.Request) {
	if Accept(r, "text/html") {
		a.HTMLReader().Read(w, a.RootType, r.URL.Path)
	}
}
