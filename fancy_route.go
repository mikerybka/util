package util

import (
	"context"
	"net/http"
)

type FancyRoute struct {
	Root   http.Handler
	Static map[string]http.Handler

	Catchall *FancyRoute
	VarName  string
}

func (route *FancyRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, isRoot := PopPath(r.URL.Path)
	if isRoot {
		route.Root.ServeHTTP(w, r)
		return
	}

	r.URL.Path = rest

	h, found := route.Static[first]
	if !found {
		if route.Catchall != nil {
			ctx := r.Context()
			params := ctx.Value("params").(map[string]string)
			params[route.VarName] = first
			ctx = context.WithValue(ctx, "params", params)
			r = r.WithContext(ctx)
			h = route.Catchall
		} else {
			h = http.NotFoundHandler()
		}
	}

	h.ServeHTTP(w, r)
}
