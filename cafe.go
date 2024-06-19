package util

import "net/http"

type Cafe[T http.Handler] struct {
	ID   string
	Data map[string]T
}
