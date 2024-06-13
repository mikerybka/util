package util

import "net/http"

type Cafe[T http.Handler] struct {
	Data map[string]T
}

func (c *Cafe[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Map[T](c.Data).ServeHTTP(w, r)
}
