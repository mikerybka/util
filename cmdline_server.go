package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

// CmdlineServer let's you easily turn any http.Handler into a cmdline app that operates on local data.
func CmdlineServer(h http.Handler) {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <filepath> <method> <url> [header1=val1,...]\nStdin is the request body.\n", os.Args[0])
		return
	}
	path := os.Args[1]
	b, _ := os.ReadFile(path)
	json.Unmarshal(b, h)
	method := os.Args[2]
	url := os.Args[3]
	headers := os.Args[4:]
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
	if IsMutation(r) {
		err = WriteJSONFile(path, h)
		if err != nil {
			panic(err)
		}
	}
	res := w.Result()
	fmt.Println(res.StatusCode)
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		panic(err)
	}
}
