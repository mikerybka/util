package util

import (
	"net/http"
	"sync"
)

func NewSingleFileServer[T http.Handler](path string) *SingleFileServer[T] {
	var v T
	ReadJSONFile(path, v)
	return &SingleFileServer[T]{
		Path: path,
		V:    v,
	}
}

type SingleFileServer[T http.Handler] struct {
	Path      string
	V         T
	writeLock sync.Mutex
}

func (s *SingleFileServer[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.V.ServeHTTP(w, r)
	if IsMutation(r) {
		go s.write()
	}
}

func (s *SingleFileServer[T]) write() {
	s.writeLock.Lock()
	defer s.writeLock.Unlock()
	err := WriteJSONFile(s.Path, s.V)
	if err != nil {
		panic("unable to write data file")
	}
}
