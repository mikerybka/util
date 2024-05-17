package util

import (
	"encoding/json"
	"net/http"
)

type SchemaCafeOrg struct {
	ID      string
	Name    string
	Schemas map[string]*Schema
}

func (org *SchemaCafeOrg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)

	if len(path) == 0 {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(org)
			return
		}

		if r.Method == "POST" {
			id := RandomID()
			schema := &Schema{
				ID: id,
			}
			json.NewDecoder(r.Body).Decode(schema)
			org.Schemas[id] = schema
			return
		}
	}

	schemaID := path[0]
	schema, ok := org.Schemas[schemaID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.StripPrefix("/"+schemaID, schema).ServeHTTP(w, r)
}
