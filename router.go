package util

import (
	"net/http"
	"slices"
	"strings"
)

type Router struct {
	Root func(w http.ResponseWriter, r *http.Request)
	Next func(first string, w http.ResponseWriter, r *http.Request)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, isRoot := PopPath(r.URL.Path)
	if isRoot {
		router.Root(w, r)
		return
	}
	r.URL.Path = rest
	router.Next(first, w, r)
}

// FancyRouter handles HTTP requests by matching each request to a handler
// according to Next.js routing conventions.
type FancyRouter struct {
	// Routes. Each key should start with '/'.
	// To handle the root request, use the key "/".
	// To define a path variable, use square brackets like "/items/[itemID]".
	Routes   map[string]http.Handler
	NotFound http.Handler
}

func (r *FancyRouter) SortedRoutes() []string {
	routes := []string{}
	for k := range r.Routes {
		routes = append(routes, k)
	}
	slices.Sort(routes)
	return routes
}

// func (r *Router) FindRoute(path string) (string, bool) {
// 	for _, route := range r.SortedRoutes() {
// 		if route == path {
// 			return route, true
// 		}

// 	}
// 	// Check for an exact match.
// 	_, ok := r.Routes[path]
// 	if ok {
// 		return path, true
// 	}

// 	// Break the path into parts
// 	pathParts := ParsePath(path)

// 	// Keep track of the possible paths
// 	possiblePaths := maps.Clone(r.Routes)

// 	// Loop through each part of the path.
// 	for i, p := range pathParts {

// 		// If there is an exact match
// 		exact := true
// 		for k := range possiblePaths {
// 			path := ParsePath(k)
// 			isVar := strings.HasPrefix(path[i], "[") && strings.HasSuffix(path[i], "]")
// 			if isVar && exact == true {
// 				exact = false
// 				continue
// 			}
// 			if isVar {
// 				delete(possiblePaths, k)
// 			}
// 			if path[i] != p && !isVar {
// 				delete(possiblePaths, k)
// 			}
// 		}
// 	}

// 	if len(possiblePaths) == 0 {
// 		return "", false
// 	}

// 	for p := range possiblePaths {
// 		if p == "" {

// 		}
// 	}

// 	if len(possiblePaths) > 1 {
// 		// TODO: handle multiple matches
// 		panic("wip")
// 	}

// 	return possiblePaths[0]
// }

func (r *FancyRouter) match(path string) (http.Handler, map[string]string) {
	pathVars := map[string]string{}

	// Check for exact match.
	h, ok := r.Routes[path]
	if ok {
		return h, pathVars
	}

	matches := []string{}
	for k := range r.Routes {
		isMatch, vars := checkMatch(k, path)
		if isMatch {
			matches = append(matches, k)
			for k, v := range vars {
				pathVars[k] = v
			}
		}
	}

	if len(matches) == 0 {
		return r.NotFound, pathVars
	}

	if len(matches) == 1 {
		return r.Routes[matches[0]], pathVars
	}

	panic("")
}

func selectMatch(matches []string) string {
	best := matches[0]
	numParts := len(ParsePath(best))
	for i := 0; i < numParts; {
	}
	return ""
}

func checkMatch(route, path string) (bool, map[string]string) {
	// Parse both the route and the path given into an array of strings so
	// that it is easier to work with.
	routeParts := ParsePath(route)
	pathParts := ParsePath(path)

	vars := map[string]string{}

	// If the number of parts in the route don't match the number of parts in
	// the path, we know it can't be a match.
	if len(routeParts) != len(pathParts) {
		return false, nil
	}

	// Loop through each part.
	for i := 0; i < len(routeParts); i++ {
		if isPathVar(routeParts[i]) {
			varName := strings.TrimSuffix(strings.TrimPrefix(path, "["), "]")
			vars[varName] = pathParts[i]
			continue
		}
		if pathParts[i] != routeParts[i] {
			return false, nil
		}
	}

	return true, vars
}

func isPathVar(s string) bool {
	return strings.HasPrefix(s, "[") && strings.HasPrefix(s, "]")
}
