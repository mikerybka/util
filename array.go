package util

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Array[T any] struct {
	Path  []string
	Value []T
}

func (a *Array[T]) ID() string {
	return JoinPath(a.Path)
}

func (a *Array[T]) JSON() string {
	b, err := json.Marshal(a.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (a *Array[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if Accept(r, "text/html") {
			fmt.Fprintf(w, "<div class='array' id='%s' data-value='%s' />", a.ID(), html.EscapeString(a.JSON()))
			return
		}

		json.NewEncoder(w).Encode(a.Value)
		return
	}

	if r.Method == "PUT" {
		err := json.NewDecoder(r.Body).Decode(a.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(a.Value)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
