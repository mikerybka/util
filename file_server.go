package util

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileServer struct {
	Root string
}

func (fs *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (fs *FileServer) get(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(fs.Root, r.URL.Path)
	fi, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "not found", http.StatusNotFound)
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
				f.Value.Files = append(f.Value.Files, e.Name())
			}
		}
	} else {
		f, err := os.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.Copy(w, f)
	}
}

func (fs *FileServer) put(w http.ResponseWriter, r *http.Request) {

}

func (fs *FileServer) delete(w http.ResponseWriter, r *http.Request) {

}

func (fs *FileServer) post(w http.ResponseWriter, r *http.Request) {

}
