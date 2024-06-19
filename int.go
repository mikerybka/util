package util

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Int struct {
	Path  []string
	Value int64
}

func (i *Int) ID() string {
	return JoinPath(i.Path)
}

func (i *Int) JSON() string {
	b, err := json.Marshal(i.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (i *Int) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if Accept(r, "text/html") {
			fmt.Fprintf(w, "<div class='int' id='%s' data-value='%s' />", i.ID(), html.EscapeString(i.JSON()))
			return
		}

		json.NewEncoder(w).Encode(i.Value)
		return
	}

	if r.Method == "PUT" {
		err := json.NewDecoder(r.Body).Decode(i.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(i.Value)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
