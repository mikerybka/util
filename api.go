package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// The API type represents an API backed by any JSON-serializable
// Go object.
type API struct {
	Type     *Type
	DataPath string
}

// ServeHTTP serves a generic REST API.
func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse the path into sections
	path := ParsePath(r.URL.Path)

	// If we're at the root, we serve it.
	if len(path) == 0 {
		api.ServeRoot(w, r)
		return
	}

	// If not, we dig one level deeper.
	first := path[0]
	rest := "/" + strings.Join(path[1:], "/")

	// If the type doesn't have nesting, return not found.
	if api.Type.IsScalar {
		api.ServeNotFound(w, r)
		return
	}

	// If it's a pointer, same thing.
	if api.Type.IsPointer {
		api.ServeNotFound(w, r)
		return
	}

	// If the data is an array
	if api.Type.IsArray {
		// As a special case, handle POST requests by appending to the array.
		if rest == "/" && r.Method == "POST" {
			// Add to the end of the array
			api.ServePOST(w, r)
			return
		}

		// Otherwise, dig one level deeper.
		next := &API{
			Type:     api.Type.ElemType,
			DataPath: filepath.Join(api.DataPath, first),
		}
		r.URL.Path = rest
		next.ServeHTTP(w, r)
		return
	}

	// If the data is a map, dig.
	if api.Type.IsMap {
		next := &API{
			Type:     api.Type.ElemType,
			DataPath: filepath.Join(api.DataPath, first),
		}
		r.URL.Path = rest
		next.ServeHTTP(w, r)
		return
	}

	if api.Type.IsStruct {
		for _, f := range api.Type.Fields {
			if f.ID == first {
				next := &API{
					Type:     api.Type.ElemType,
					DataPath: filepath.Join(api.DataPath, first),
				}
				r.URL.Path = rest
				next.ServeHTTP(w, r)
				return
			}
		}

		// If the field is not found, return not found.
		api.ServeNotFound(w, r)
	}

	panic("invalid type")
}

func (api *API) ServeNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("null"))
}

func (api *API) ServePOST(w http.ResponseWriter, r *http.Request) {
	// Get the size of the dir
	entries, err := os.ReadDir(api.DataPath)
	if err != nil {
		panic(err)
	}
	dirSize := len(entries)

	// Figure out the next ID
	nextID := strconv.Itoa(dirSize)

	// Read the data
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Create a new object with that ID
	err = WriteFile(filepath.Join(api.DataPath, nextID), b)
	if err != nil {
		panic(err)
	}
}

func (api *API) ServeRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// If the value is a scalar or a pointer, read from file.
		if api.Type.IsScalar || api.Type.IsPointer {
			b, err := os.ReadFile(api.DataPath)
			if err != nil {
				panic(err)
			}
			w.Write(b)
			return
		}

		// If the value is an array or a map, read the directory.
		if api.Type.IsArray || api.Type.IsMap {
			entries, err := os.ReadDir(api.DataPath)
			if err != nil {
				panic(err)
			}
			items := []string{}
			for _, entry := range entries {
				items = append(items, entry.Name())
			}
			return
		}

		// If the value is a struct, build a JSON object.
		if api.Type.IsStruct {
			// Open the object.
			w.Write([]byte("{"))

			// Write each field out.
			for i, f := range api.Type.Fields {
				// Write the field ID as the key.
				fmt.Fprintf(w, "\"%s\": ", f.ID)

				// Build a new API object and serve the field.
				fieldAPI := &API{
					Type:     f.Type,
					DataPath: filepath.Join(api.DataPath, f.ID),
				}
				fieldAPI.ServeHTTP(w, r)

				// Add a comma if it's not the last field.
				if i != len(api.Type.Fields)-1 {
					fmt.Fprintf(w, ",")
				}
			}

			// Close object.
			w.Write([]byte("}"))
			return
		}

		panic("invalid type")
	}

	if r.Method == "PUT" {
		// If the value is a scalar or a pointer, write to file.
		if api.Type.IsScalar || api.Type.IsPointer {

			return
		}

		// Otherwise, don't allow the request.
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
