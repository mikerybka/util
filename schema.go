package util

import (
	"net/http"
)

type Schema struct {
	ID     string
	Path   string
	Name   string
	Fields []Field
}

func (s *Schema) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st := &Struct[*Schema]{
		Path: s.Path,
		Data: s,
	}

	st.ServeHTTP(w, r)
}
