package main

import (
	"net/http"

	"github.com/mikerybka/util"
)

func main() {
	app := &util.Cafe[*util.Schema]{
		Data: util.Map[*util.Schema]{
			"tuv": &util.Schema{
				ID:     "tuv",
				Name:   "TUV",
				Fields: []util.Field{},
			},
		},
	}
	err := http.ListenAndServe(":8000", app)
	panic(err)
}
