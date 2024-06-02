package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type BuildServer struct {
	Workdir string
}

func (s *BuildServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
