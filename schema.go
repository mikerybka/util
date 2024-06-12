package util

import (
	"fmt"
	"net/http"
)

type Schema struct {
	ID     string
	Path   string
	Name   string
	Fields []Field
}

func (s *Schema) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div class=\"schema\">")
	f := &Field{
		ID:   "id",
		Name: "ID",
		Type: "string",
	}
	f.ServeHTTP(w, r)
	f = &Field{
		ID:   "fields",
		Name: "Fields",
		Type: "[]util.Field",
	}
	f.ServeHTTP(w, r)
	fmt.Fprintf(w, "</div>")
}
