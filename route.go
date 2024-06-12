package util

import "net/http"

type Route struct {
	Root    http.Handler
	Static  map[string]http.Handler
	Dynamic func(string) http.Handler
}

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, isRoot := PopPath(r.URL.Path)
	if isRoot {
		route.Root.ServeHTTP(w, r)
		return
	}

	r.URL.Path = rest

	h, found := route.Static[first]
	if !found {
		if route.Dynamic != nil {
			h = route.Dynamic(first)
		} else {
			h = http.NotFoundHandler()
		}
	}

	h.ServeHTTP(w, r)
}
