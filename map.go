package util

import (
	"fmt"
	"net/http"
)

type Map[T http.Handler] map[string]T

func (m Map[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := &Route{
		Root: MapList[T](m),
		Dynamic: func(s string) http.Handler {
			fmt.Println(s)
			fmt.Println(m)
			h, ok := m[s]
			fmt.Println(ok)
			if !ok {
				return http.NotFoundHandler()
			}
			return h
		},
	}
	route.ServeHTTP(w, r)
}
