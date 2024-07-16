package util

import "net/http"

type Route struct {
	GetRoot  http.HandlerFunc
	PostRoot http.HandlerFunc
	GetID    func(id string, w http.ResponseWriter, r *http.Request)
	PostID   func(id string, w http.ResponseWriter, r *http.Request)
	PutID    func(id string, w http.ResponseWriter, r *http.Request)
	PatchID  func(id string, w http.ResponseWriter, r *http.Request)
	DeleteID func(id string, w http.ResponseWriter, r *http.Request)
}

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, rest, isRoot := PopPath(r.URL.Path)
	if isRoot {
		switch r.Method {
		case http.MethodGet:
			if route.GetRoot == nil {
				http.NotFound(w, r)
			} else {
				route.GetRoot.ServeHTTP(w, r)
			}
		case http.MethodPost:
			if route.PostRoot == nil {
				http.NotFound(w, r)
			} else {
				route.PostRoot.ServeHTTP(w, r)
			}
		}
	} else {
		r.URL.Path = rest
		switch r.Method {
		case http.MethodGet:
			if route.GetID == nil {
				http.NotFound(w, r)
			} else {
				route.GetID(id, w, r)
			}
		case http.MethodPost:
			if route.PostID == nil {
				http.NotFound(w, r)
			} else {
				route.PostID(id, w, r)
			}
		case http.MethodPut:
			if route.PutID == nil {
				http.NotFound(w, r)
			} else {
				route.PutID(id, w, r)
			}
		case http.MethodPatch:
			if route.PatchID == nil {
				http.NotFound(w, r)
			} else {
				route.PatchID(id, w, r)
			}
		case http.MethodDelete:
			if route.DeleteID == nil {
				http.NotFound(w, r)
			} else {
				route.DeleteID(id, w, r)
			}
		}
	}
}
