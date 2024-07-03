package util

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func NewRequest(r *http.Request) *Request {
	body, _ := io.ReadAll(r.Body)
	copy := io.NopCloser(bytes.NewReader(body))
	r.Body = copy
	return &Request{
		Method: r.Method,
		URL:    r.URL.String(),
		Header: r.Header,
		Body:   string(body),
	}
}

type Request struct {
	Method string
	URL    string
	Header http.Header
	Body   string
}

func (r *Request) Log() {
	b, _ := json.Marshal(r)
	log.Println(string(b))
}
