package serverelfect

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/mikerybka/util"
)

func Serve(v reflect.Value, w http.ResponseWriter, r *http.Request) {
	path := util.ParsePath(r.URL.Path)
	if len(path) == 0 {
		serveRoot(v, w, r)
	}
	first, path := util.Pop(path)
	r.URL.Path = "/" + strings.Join(path, "/")
	if isMethod(v, first) {
		serveMethod(v, first, w, r)
	} else if isField(v, first) {
		serveField(v, first, w, r)
	} else if isItem(v, first) {
		serveItem(v, first, w, r)
	} else {
		serveNull(w, r)
	}
}

func serveNull(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func serveRoot(v reflect.Value, w http.ResponseWriter, r *http.Request)
func isMethod(v reflect.Value, name string) bool
func serveMethod(v reflect.Value, name string, w http.ResponseWriter, r *http.Request)
func isField(v reflect.Value, name string) bool
func serveField(v reflect.Value, name string, w http.ResponseWriter, r *http.Request)
func isItem(v reflect.Value, name string) bool
func serveItem(v reflect.Value, name string, w http.ResponseWriter, r *http.Request)
