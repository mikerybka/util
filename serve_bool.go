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
			if v.(bool) {
				fmt.Fprintf(w, "<input type=\"checkbox\" id=\"%s\" checked />", "/"+strings.Join(path, "/"))
				return
			} else {
				fmt.Fprintf(w, "<input type=\"checkbox\" id=\"%s\" />", "/"+strings.Join(path, "/"))
				return
			}
		}

		json.NewEncoder(w).Encode(v)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
