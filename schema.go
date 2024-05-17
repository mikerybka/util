package util

import (
	"encoding/json"
	"net/http"
)

type Schema struct {
	ID     string
	Name   string
	Fields []Field
}

func (s *Schema) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(s)
		return
	}

	if r.Method == "PATCH" {
		json.NewDecoder(r.Body).Decode(s)
		return
	}
}
