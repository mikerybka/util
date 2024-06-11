package main

import (
	"net/http"

	"github.com/mikerybka/util"
)

func main() {
	app := &util.Cafe[util.Schema]{}
	err := http.ListenAndServe(":8000", app)
	panic(err)
}
