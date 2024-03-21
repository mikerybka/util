package util

import (
	"net/http"
	"os"
)

type API[T any] struct {
	Data T
}

func (a *API[T]) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return http.ListenAndServe(":"+port, a)
}

func (api *API[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if api == nil {
		WriteNotFound(w)
		return
	}
	path := ParsePath(r.URL.Path)
	switch len(path) {
	case 0:
		switch r.Method {
		case "GET":
			WriteJSON(w, api.Data)
		case "POST":
			if IsArray(api.Data) {
				panic("not implemented")
			} else if IsMap(api.Data) {
				panic("not implemented")
			} else {
				WriteMethodNotAllowed(w)
			}
		case "PUT":
			HandlePUT(w, r, api.Data)
		default:
			WriteNotFound(w)
		}
	case 1:
		if IsArray(api.Data) {
			switch r.Method {
			case "DELETE":
				panic("not implemented")
			case "PUT":
				panic("not implemented")
			default:
				panic("not implemented")
			}
		} else if IsMap(api.Data) {
			switch r.Method {
			case "DELETE":
				panic("not implemented")
			case "PUT":
				panic("not implemented")
			default:
				panic("not implemented")
			}
		} else {
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}
