package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mikerybka/util"
)

func main() {
	s := &util.Server{}
	f, _ := os.Open("dev/data.json")
	json.NewDecoder(f).Decode(s)
	panic(http.ListenAndServe(":8000", s))
}
