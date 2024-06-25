package util

import "net/http"

type SchemaCafe struct {
	Schemas map[string]*Schema
}

func (sc *SchemaCafe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := []string{}
	value := map[string]any{}
	for k, v := range sc.Schemas {
		value[k] = v
	}
	obj := &Map{
		Path:  path,
		Value: value,
	}
	ServeObject(path, obj, w, r)
}
