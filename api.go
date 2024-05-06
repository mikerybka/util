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

func (api *API[T]) handleArrayPOST(w http.ResponseWriter, r *http.Request) {
	slice := reflect.ValueOf(api.Data)
	newSlice := reflect.MakeSlice(slice.Type(), slice.Len()+1, slice.Cap()+1)
	reflect.Copy(newSlice, slice)

	v := reflect.New(slice.Type().Elem())
	vP := v.Interface()
	err := json.NewDecoder(r.Body).Decode(&vP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slice = reflect.Append(slice, reflect.ValueOf(v))
	api.Data = slice.Interface().(T)
	json.NewEncoder(w).Encode(v)
}

func (api *API[T]) handleArrayPUT(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)
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
	json.NewEncoder(w).Encode(newValue)
}

func (api *API[T]) handleMapPUT(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)
	v := reflect.ValueOf(api.Data)
	key := reflect.ValueOf(path[0])
	newValue := reflect.New(v.Type().Elem())
	a := newValue.Interface()
	err := json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	v.SetMapIndex(key, newValue)
	json.NewEncoder(w).Encode(a)
}

func (api *API[T]) handleArrayDELETE(w http.ResponseWriter, r *http.Request) {
	idx, err := strconv.Atoi(ParsePath(r.URL.Path)[0])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	arr := reflect.ValueOf(api.Data)
	if idx < 0 || idx >= arr.Len() {
		http.NotFound(w, r)
		return
	}
	newArr := reflect.AppendSlice(arr.Slice(0, idx), arr.Slice(idx+1, arr.Len()))
	api.Data = newArr.Interface().(T)
}

func (api *API[T]) handleMapDELETE(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (api *API[T]) handleStructPUT(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)
	// https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
	v := reflect.ValueOf(api.Data)
	f := v.FieldByName(path[0])
	newValue := reflect.New(f.Addr().Type())
	a := newValue.Interface()
	err := json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f.Set(newValue)
}

func (api *API[T]) hasMethod(m string) bool {
	v := reflect.ValueOf(api.Data)
	meth := v.MethodByName(m)
	return !meth.IsZero()
}

func (api *API[T]) handleMethodCall(m string, w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
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
				api.handleArrayPOST(w, r)
			} else {
				WriteMethodNotAllowed(w)
			}
		case "PUT":
			HandlePUT(w, r, &api.Data)
		default:
			WriteNotFound(w)
		}
	} else if len(path) == 1 {
		if r.Method == "PUT" {
			if IsArray(api.Data) {
				api.handleArrayPUT(w, r)
			} else if IsMap(api.Data) {
				api.handleMapPUT(w, r)
			} else if IsStruct(api.Data) {
				api.handleStructPUT(w, r)
			}
		} else if r.Method == "DELETE" {
			if IsArray(api.Data) {
				api.handleArrayDELETE(w, r)
			} else if IsMap(api.Data) {
				api.handleMapDELETE(w, r)
			}
		} else if api.hasMethod(path[0]) {
			api.handleMethodCall(path[0], w, r)
		}
	} else {
		r.URL.Path = "/" + strings.Join(path[1:], "/")
		api.dig(path[0]).ServeHTTP(w, r)
	}
}
