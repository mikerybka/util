package util

import (
	"encoding/json"
	"net/http"
)

type JSONFile[T http.Handler] struct {
	Path string
}

func (f *JSONFile[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var v T
	ReadJSONFile(f.Path, v)
	v.ServeHTTP(w, r)
	if IsMutation(r) {
		b, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		err = WriteFile(f.Path, b)
		if err != nil {
			panic(err)
		}
	}
}
