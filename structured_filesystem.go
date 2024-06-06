package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type StructuredFilesystem struct {
	Types    map[string]Type
	RootType string
	Data     FileSystem
}

func (fs *StructuredFilesystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fs.ServeGET(w, r)
	case http.MethodPost:
		fs.ServePOST(w, r)
	case http.MethodPut:
		fs.ServePUT(w, r)
	case http.MethodDelete:
		fs.ServeDELETE(w, r)
	case http.MethodOptions:
		fs.ServeOPTIONS(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (fs *StructuredFilesystem) Get(path string) ([]byte, error) {
	t, ok := fs.Types[fs.RootType]
	if !ok {
		panic(fmt.Errorf("unkown type: %s", fs.RootType))
	}

	first, rest, ok := PopPath(path)
	isRoot := !ok
	if !isRoot {
		if t.IsScalar || t.IsPointer {
			return nil, os.ErrNotExist
		}
		if t.IsArray || t.IsMap {
			subfs := &StructuredFilesystem{
				Types:    fs.Types,
				RootType: t.ElemType,
				Data:     fs.Data.Dig(first),
			}
			return subfs.Get(rest)
		}
		if t.IsStruct {
			for _, f := range t.Fields {
				if f.ID == first {
					subfs := &StructuredFilesystem{
						Types:    fs.Types,
						RootType: f.Type,
						Data:     fs.Data.Dig(first),
					}
					return subfs.Get(rest)
				}
			}
			return nil, os.ErrNotExist
		}
		panic("invalid type")
	} else {
		if t.IsScalar || t.IsPointer {
			return fs.Data.ReadFile("/")
		}
		if t.IsArray || t.IsMap || t.IsStruct {
			entries, err := fs.Data.ReadDir("/")
			if err != nil {
				return nil, err
			}
			b, err := json.Marshal(entries)
			if err != nil {
				panic(err)
			}
			return b, nil
		}
		panic("invalid type")
	}
}

func (fs *StructuredFilesystem) ServeGET(w http.ResponseWriter, r *http.Request) {
	b, err := fs.Get(r.URL.Path)
	if errors.Is(err, os.ErrNotExist) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func (fs *StructuredFilesystem) ServePOST(w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}
func (fs *StructuredFilesystem) ServePUT(w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}
func (fs *StructuredFilesystem) ServeDELETE(w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}
func (fs *StructuredFilesystem) ServeOPTIONS(w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}
