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
	ServeStruct(w, r)
}

func ServeStruct(w http.ResponseWriter, r *http.Request) {}
