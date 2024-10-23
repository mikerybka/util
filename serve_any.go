package util

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

// ServeAny serves any Go value.
// Drill into any exported struct field or map/arary value.
// Arrays start at 1001.
// GET requests are served the JSON encoded values.
// POST requests are used to call methods.
// Ex request body: `{"Method":"Add","Args":["1","2","3"]}`.
// All args are JSON-encoded strings.
// The server will unmarshal each arg into it's appropriate type.
// Ex response body: `["6","null"]`.
func ServeAny(v any, w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path[1:], "/")
	v, _ = Drill(v, path)

	if r.Method == "GET" {
		val := struct {
			Metadata Metadata `json:"metadata"`
			Data     any      `json:"data"`
		}{
			Metadata: Metadata{
				Type: reflect.TypeOf(v).String(),
			},
			Data: v,
		}
		json.NewEncoder(w).Encode(&val)
		return
	}

	if r.Method == "POST" {
		// Parse
		req := ReadJSON[MethodCall](r.Body)

		// Execute
		res, err := CallMethod(v, req.Method, req.Args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond
		json.NewEncoder(w).Encode(res)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
