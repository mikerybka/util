package util

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func NewStreamMultiplexer(contentType string, backendURL string) *StreamMultiplexer {
	return &StreamMultiplexer{
		proxy: NewReverseProxy(backendURL),
		broadcaster: NewStreamBroadcaster(contentType, func() (io.ReadCloser, error) {
			req, err := http.NewRequest("GET", backendURL, nil)
			if err != nil {
				return nil, err
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return nil, err
			}
			if res.StatusCode != 200 {
				return nil, fmt.Errorf("%s", res.Status)
			}
			return res.Body, nil
		}),
	}
}

type StreamMultiplexer struct {
	broadcaster *StreamBroadcaster
	proxy       *httputil.ReverseProxy
}

func (m *StreamMultiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		m.broadcaster.ServeHTTP(w, r)
	} else {
		m.proxy.ServeHTTP(w, r)
	}
}
