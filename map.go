package util

import (
	"net/http"
)

type Map[T http.Handler] struct {
	ID   string
	Data map[string]T
}

func (m *Map[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := &FancyRoute{
		Root: &MapList[T]{
			ID:   m.ID,
			Data: m.Data,
		},
		Catchall: nil,
	}
	route.ServeHTTP(w, r)
}
