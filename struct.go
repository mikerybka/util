package util

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
)

type Struct[T any] struct {
	Path string
	Data T
}

func (s *Struct[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := reflect.TypeOf(s.Data)
	for {
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		} else {
			break
		}
	}
	if t.Kind() != reflect.Struct {
		panic("expected struct")
	}
	fmt.Fprintf(w, "<div class=\"struct\">")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Fprintf(w, "<div class=\"field\"><a href=\"%s\">%s</a></div>",
			filepath.Join(r.URL.Path, f.Name), f.Name)
	}
	fmt.Fprintf(w, "</div>")
}
