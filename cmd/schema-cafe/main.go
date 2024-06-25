package main

import (
	"net/http"

	"github.com/mikerybka/util"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := []string{}
		obj := &util.Map{
			Path: path,
			Value: map[string]any{
				"schema1": &util.Schema{
					Fields: []util.SchemaField{
						{
							Name: "yo",
							Type: "string",
						},
					},
				},
			},
		}
		util.ServeObject(path, obj, w, r)
	})
	err := http.ListenAndServe(":8000", nil)
	panic(err)
}
