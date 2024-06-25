package util

import (
	"encoding/json"
	"net/http"
	"os"
)

type FileObject[T Object] struct {
	Path string
}

func (fo *FileObject[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var v T
	f, _ := os.Open(fo.Path)
	json.NewDecoder(f).Decode(v)
	f.Close()

	ServeObject([]string{}, v, w, r)

	if IsMutation(r) {
		b, _ := json.Marshal(v)
		err := WriteFile(fo.Path, b)
		if err != nil {
			panic(err)
		}
	}
}
