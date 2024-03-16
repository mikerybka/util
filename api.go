package util

import (
	"encoding/json"
	"net/http"
)

type API[T any] struct {
	Data T
}

func (api *API[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)
	if len(path) == 0 {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(api.Data)
			return
		}
		if r.Method == "POST" {
			// if T is an array {
			// TODO: handle adding to it
			// }
			return
		}
		if r.Method == "PUT" {
			err := json.NewDecoder(r.Body).Decode(&api.Data)
			if err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			return
		}
	}
	if len(path) == 1 {
		if r.Method == "POST" {
			// TODO: handle method calls.
		}
	}
	// TODO: drill into maps arrays and structs
}
