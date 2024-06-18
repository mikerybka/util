package util

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

// Router struct
type Router struct {
	routes []*Route
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{}
}

// AddRoute adds a new route to the router
func (r *Router) AddRoute(path string, handler http.HandlerFunc) {
	// Convert path with variables to regex
	regexPattern := "^" + path
	regexPattern = strings.Replace(regexPattern, "[", "(?P<", -1)
	regexPattern = strings.Replace(regexPattern, "]", ">[^/]+)", -1)
	regexPattern += "$"
	regex := regexp.MustCompile(regexPattern)

	r.routes = append(r.routes, &Route{
		Pattern: regex,
		Handler: handler,
	})
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if match := route.Pattern.FindStringSubmatch(req.URL.Path); match != nil {
			// Extract the path variables
			params := make(map[string]string)
			for i, name := range route.Pattern.SubexpNames() {
				if i > 0 && name != "" {
					params[name] = match[i]
				}
			}
			// Attach params to the request context
			ctx := req.Context()
			ctx = context.WithValue(ctx, "params", params)
			req = req.WithContext(ctx)

			// Call the handler
			route.Handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// GetParam retrieves the path variable from the request context
func GetParam(r *http.Request, name string) string {
	params, ok := r.Context().Value("params").(map[string]string)
	if !ok {
		return ""
	}
	return params[name]
}
