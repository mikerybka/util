package util

import "net/http"

type Cafe[T http.Handler] struct {
	ID   string
	Data map[string]T
}

func (c *Cafe[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := &Map[T]{
		ID:   c.ID,
		Data: c.Data,
	}
	m.ServeHTTP(w, r)
}
