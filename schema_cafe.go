package util

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type SchemaCafe struct {
	Schemas map[string]*Schema
}

func (s *SchemaCafe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)

	if len(path) == 0 {
		fmt.Fprintf(w, "<div class=\"schema-list\">")
		for id := range s.Schemas {
			fmt.Fprintf(w, "<a href=\"%s\">%s</a>", filepath.Join(r.URL.Path, id), id)
		}
		fmt.Fprintf(w, "</div>")
		return
	}

	schemaID := path[0]
	schema, ok := s.Schemas[schemaID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	schema.ServeHTTP(w, r)
}
