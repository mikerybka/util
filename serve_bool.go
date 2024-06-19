package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ServeBool(path []string, v any, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if Accept(r, "text/html") {
			fmt.Fprintf(w, "<input type=\"checkbox\" id=\"%s\">%d</input", "/"+strings.Join(path, "/"), v.(bool))
			return
		}

		json.NewEncoder(w).Encode(v)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
