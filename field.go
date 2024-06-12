package util

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type Field struct {
	ID   string
	Name string
	Type string
}

func (f *Field) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := f.Name
	if name == "" {
		name = f.ID
	}
	fmt.Println(r.URL.Path)
	l := &Link{
		Name: name,
		Href: filepath.Join(r.URL.Path, f.ID),
	}
	l.ServeHTTP(w, r)
}
