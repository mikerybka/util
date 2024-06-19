package util

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Bool struct {
	Path  []string
	Value bool
}

func (b *Bool) ID() string {
	return JoinPath(b.Path)
}

func (b *Bool) JSON() string {
	byt, err := json.Marshal(b.Value)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

func (b *Bool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if Accept(r, "text/html") {
			fmt.Fprintf(w, "<div class='bool' id='%s' data-value='%s' />", b.ID(), html.EscapeString(b.JSON()))
			return
		}

		json.NewEncoder(w).Encode(b.Value)
		return
	}

	if r.Method == "PUT" {
		err := json.NewDecoder(r.Body).Decode(b.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(b.Value)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
