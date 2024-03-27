package util

import (
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// NewAPI creates a new API object.
func NewAPI[T any](d T) *API[T] {
	return &API[T]{
		Data: d,
	}
}

// The API type represents an API backed by any JSON-serializable
// Go object.
type API[T any] struct {
	Data T
}

// Start binds to PORT (9000 by default) and handles API requests.
func (a *API[T]) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return http.ListenAndServe(":"+port, a)
}

func (a *API[T]) dig(n string) *API[any] {
	v := reflect.ValueOf(a.Data)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		idx, err := strconv.Atoi(n)
		if err != nil {
			return nil
		}
		if idx < 0 || idx >= v.Len() {
			return nil
		}
		item := v.Index(idx)
		return &API[any]{
			Data: item.Interface(),
		}
	case reflect.Map:
		key := reflect.ValueOf(n)
		item := v.MapIndex(key)
		return &API[any]{
			Data: item.Interface(),
		}
	case reflect.Struct:
		field := v.FieldByName(n)
		return &API[any]{
			Data: field.Interface(),
		}
	default:
		return nil
	}
}

// ServeHTTP serves a generic REST API.
func (api *API[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if api == nil {
		WriteNotFound(w)
		return
	}
	path := ParsePath(r.URL.Path)
	if len(path) == 0 {
		switch r.Method {
		case "GET":
			WriteJSON(w, api.Data)
		case "POST":
			if IsArray(api.Data) {
				panic("not implemented")
			} else {
				WriteMethodNotAllowed(w)
			}
		case "PUT":
			HandlePUT(w, r, &api.Data)
		default:
			WriteNotFound(w)
		}
		return
	}
	if len(path) == 1 {
		if r.Method == "PUT" {
			if IsArray(api.Data) {
				v := reflect.ValueOf(api.Data)
				idx, err := strconv.Atoi(path[0])
				if err != nil {
					WriteNotFound(w)
					return
				}
				if idx < 0 || idx >= v.Len() {
					WriteNotFound(w)
					return
				}
				item := v.Index(idx)
				newValue := reflect.New(item.Type()).Interface()
				err = json.NewDecoder(r.Body).Decode(newValue)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				item.Set(reflect.ValueOf(newValue))
				return
			} else if IsMap(api.Data) {
				panic("not implemented")
				return
			} else if IsStruct(api.Data) {
				panic("not implemented")
				return
			}
		}
		if r.Method == "DELETE" {
			if IsArray(api.Data) {
				panic("not implemented")
			} else if IsMap(api.Data) {
				panic("not implemented")
			}
		}
	}
	r.URL.Path = "/" + strings.Join(path[1:], "/")
	api.dig(path[0]).ServeHTTP(w, r)
}
