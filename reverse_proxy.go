package util

import (
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(backendURL string) *httputil.ReverseProxy {
	u, err := url.Parse(backendURL)
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(u)
}
