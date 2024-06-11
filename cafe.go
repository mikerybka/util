package util

import "net/http"

type Cafe[T any] struct {
	Data Map[T]
}

func (c *Cafe[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := &Route{
		Root: c.Data,
	}
}
