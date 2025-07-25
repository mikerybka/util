package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type JSONFileServer struct {
	Root string
}

func (fs *JSONFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fs.get(w, r)
	case "PUT":
		fs.put(w, r)
	case "DELETE":
		fs.delete(w, r)
	case "POST":
		fs.post(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}
}

func (fs *JSONFileServer) get(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(fs.Root, r.URL.Path) + ".json"
	fi, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			path = strings.TrimSuffix(path, ".json")
			f, err := os.Open(path)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					http.Error(w, "not found", http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			io.Copy(w, f)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if fi.IsDir() {
		f := &File[Folder]{
			ID:    r.URL.Path,
			Type:  "folder",
			Value: Folder{},
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, e := range entries {
			if e.IsDir() {
				f.Value.Folders = append(f.Value.Folders, e.Name())
			} else {
				if strings.HasSuffix(e.Name(), ".json") {
					f.Value.Files = append(f.Value.Files, strings.TrimSuffix(e.Name(), ".json"))
				}
			}
		}
		json.NewEncoder(w).Encode(f)
		return
	}
	http.NotFound(w, r)
}

func (fs *JSONFileServer) put(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(fs.Root, r.URL.Path) + ".json"
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(f, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (fs *JSONFileServer) delete(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(fs.Root, r.URL.Path)
	err := os.RemoveAll(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	path += ".json"
	err = os.RemoveAll(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (fs *JSONFileServer) post(w http.ResponseWriter, r *http.Request) {
	id := RandomString(12, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r.URL.Path = filepath.Join(r.URL.Path, id)
	fs.put(w, r)
}
