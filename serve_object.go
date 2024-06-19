package util

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
)

func ServeObject(path []string, v Object, w http.ResponseWriter, r *http.Request) {
	// Parse the request path.
	p := ParsePath(r.URL.Path)

	// Handle method requests
	if (r.Method == "POST" || r.Method == "GET") && len(p) == 1 && HasMethod(v, path[0]) {
		ServeMethod(path[0], v, w, r)
		return
	}

	// Handle DELETE requests to maps and arrays
	if r.Method == "DELETE" && len(p) == 1 && (strings.HasPrefix(v.Type(), "map[string]") || strings.HasPrefix(v.Type(), "[]")) {
		//TODO
		return
	}

	// Dig if request path is not root.
	if len(p) > 0 {
		first := p[0]
		rest := p[1:]
		r.URL.Path = JoinPath(rest)
		next, ok := v.Dig(first)
		if !ok {
			http.NotFound(w, r)
			return
		}
		ServeObject(append(path, first), next, w, r)
		return
	}

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

	// Disallow all other requests.
	http.NotFound(w, r)
	return
}
