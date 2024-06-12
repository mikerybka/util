package util

import "net/http"

type Cafe[T http.Handler] struct {
	Data Map[T]
}

func (c *Cafe[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Data.ServeHTTP(w, r)
}
