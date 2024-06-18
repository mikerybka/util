package util

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type MapList[T http.Handler] struct {
	ID   string
	Data map[string]T
}

func (m *MapList[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div class=\"list\">")
	for k, v := range m.Data {
		l := &Link{
			Name: GetName(v),
			Href: filepath.Join(m.ID, k),
		}
		l.ServeHTTP(w, r)
	}
	fmt.Fprintf(w, "</div>")
}
