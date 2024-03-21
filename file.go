package util

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

func NewFile[T any](path string) *File[T] {
	return &File[T]{
		Path: path,
	}
}

type File[T any] struct {
	Path      string
	Data      T
	WriteLock sync.Mutex
}

func (f *File[T]) Read() error {
	b, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, f.Data)
}

func (f *File[T]) Write() error {
	f.WriteLock.Lock()
	defer f.WriteLock.Unlock()
	return WriteJSONFile(f.Path, f.Data)
}

func (f *File[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	NewAPI(f.Data).ServeHTTP(w, r)

	if IsMutation(r) {
		err := f.Write()
		if err != nil {
			panic(err)
		}
	}
}
