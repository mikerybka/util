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

func (s *Schema) ReactComponent() *ReactComponent {
	return &ReactComponent{
		Name: s.Name,
		Props: []*Field{
			{
				ID:   "url",
				Name: "URL",
				Type: "string",
			},
		},
		Body: []*Statement{
			{
				IsAssign: true,
				Name:     "schema",
				Value:    &Expression{},
			},
		},
	}
}

func (s *Schema) Type() *Type {
	return &Type{
		IsStruct: true,
		Fields:   s.Fields,
	}
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
