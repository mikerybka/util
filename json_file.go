package util

import (
	"encoding/json"
	"net/http"
	"os"
)

type JSONFile[T any] struct {
	Path string
}

func (jf *JSONFile[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(jf.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var v T
	err = json.NewDecoder(f).Decode(&v)
	f.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api := &API[T]{
		Data: v,
	}

	api.ServeHTTP(w, r)

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(jf.Path, b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
