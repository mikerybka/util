package util

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

// CmdlineServer let's you easily turn any http.Handler into a cmdline app.
func CmdlineServer(h http.Handler) {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <method> <url> [header1=val1,...]\nStdin is the request body.\n", os.Args[0])
		return
	}
	method := os.Args[1]
	url := os.Args[2]
	headers := os.Args[3:]
	r, err := http.NewRequest(method, url, os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, head := range headers {
		k, v, ok := strings.Cut(head, "=")
		if !ok {
			panic("bad header: " + head)
		}
		r.Header.Add(k, v)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	res := w.Result()
	fmt.Println(res.StatusCode)
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		panic(err)
	}
}
