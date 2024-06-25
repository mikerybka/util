package util

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
)

func ServeObject(path []string, v Object, w http.ResponseWriter, r *http.Request) {
	router := &Router{
		Root: func(w http.ResponseWriter, r *http.Request) {
			// Handle POST requests to arrays
			if r.Method == "POST" && IsArray(v) {
				//TODO
				return
			}

			// Handle GET requests (JSON and HTML).
			if r.Method == "GET" {
				if Accept(r, "text/html") {
					fmt.Fprintf(w, "<div class='%s' id='%s' data-value='%s' />", v.Type(), v.ID(), html.EscapeString(v.JSON()))
					return
				}

				w.Write([]byte(v.JSON()))
				return
			}

			// Handle PUT requests (JSON body expected).
			if r.Method == "PUT" {
				err := json.NewDecoder(r.Body).Decode(v.Ptr())
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				w.Write([]byte(v.JSON()))
				return
			}
		},
		Next: func(first string, w http.ResponseWriter, r *http.Request) {
			router := &Router{
				Root: func(w http.ResponseWriter, r *http.Request) {
					// Handle method requests
					if (r.Method == "POST" || r.Method == "GET") && HasMethod(v, first) {
						ServeMethod(first, v, w, r)
						return
					}

					// Handle DELETE requests to maps and arrays
					if r.Method == "DELETE" && (strings.HasPrefix(v.Type(), "map[string]") || strings.HasPrefix(v.Type(), "[]")) {
						//TODO
						return
					}

					// Dig
					v, ok := v.Dig(first)
					if !ok {
						http.NotFound(w, r)
						return
					}
					ServeObject(append(path, first), v, w, r)
				},
				Next: func(second string, w http.ResponseWriter, r *http.Request) {
					v, ok := v.Dig(first)
					if !ok {
						http.NotFound(w, r)
						return
					}
					v, ok = v.Dig(second)
					if !ok {
						http.NotFound(w, r)
						return
					}
					ServeObject(append(path, first, second), v, w, r)
				},
			}
			router.ServeHTTP(w, r)
		},
	}

	router.ServeHTTP(w, r)
}
