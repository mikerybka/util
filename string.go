package util

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type String struct {
	Path  []string
	Value string
}

func (s *String) ID() string {
	return JoinPath(s.Path)
}

func (s *String) JSON() string {
	b, err := json.Marshal(s.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if Accept(r, "text/html") {
			fmt.Fprintf(w, "<div class='string' id='%s' data-value='%s' />", s.ID(), html.EscapeString(s.JSON()))
			return
		}

		json.NewEncoder(w).Encode(s.Value)
		return
	}

	if r.Method == "PUT" {
		err := json.NewDecoder(r.Body).Decode(s.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(s.Value)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
