package util

import (
	"encoding/json"
	"net/http"
	"os"
)

type JSONFile[T any] struct {
	Path string
}

func (f *JSONFile[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)
	if len(path) == 0 {
		if r.Method == "GET" {
			b, err := os.ReadFile(f.Path)
			if err != nil {
				panic(err)
			}
			w.Write(b)
			return
		}
		if r.Method == "PUT" {
			var v T
			err := json.NewDecoder(r.Body).Decode(&v)
			if err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(f.Path, b, os.ModePerm)
			if err != nil {
				panic(err)
			}
			return
		}
	}
	if len(path) == 1 {
		if r.Method == "POST" {
			// TODO: handle method calls.
		}
	}
}
