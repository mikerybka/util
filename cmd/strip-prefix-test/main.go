package main

import (
	"fmt"
	"net/http"
)

func main() {
	t := &Test{}
	s := http.StripPrefix("/test", t)
	http.ListenAndServe(":4000", s)
}

type Test struct {
}

func (t *Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
}
