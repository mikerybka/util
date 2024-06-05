package util

import "net/http"

type PingServer struct {
	Msg string
}

func (s *PingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.Msg == "" {
		s.Msg = "OK"
	}
	w.Write([]byte(s.Msg))
}
