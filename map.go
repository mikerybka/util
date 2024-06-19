package util

import (
	"encoding/json"
	"net/http"
)

type Map[T http.Handler] struct {
	Path  []string
	Value map[string]T
}

func (m *Map[T]) ID() string {
	return JoinPath(m.Path)
}

func (m *Map[T]) JSON() string {
	b, err := json.Marshal(m.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
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
