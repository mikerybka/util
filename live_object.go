package util

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

type LiveObject[T http.Handler] struct {
	Path  string
	Value T
	mu    sync.Mutex
}

func (obj *LiveObject[T]) Load() {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	b, _ := os.ReadFile(obj.Path)
	json.Unmarshal(b, obj.Value)
}

func NewLiveObject[T http.Handler](path string, value T) *LiveObject[T] {
	obj := &LiveObject[T]{
		Path:  path,
		Value: value,
	}
	OnFileChange(path, obj.Load)
	return obj
}

func (obj *LiveObject[T]) Save() error {
	return WriteJSONFile(obj.Path, obj.Value)
}

func (obj *LiveObject[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	obj.Value.ServeHTTP(w, r)
	if IsMutation(r) {
		err := obj.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
