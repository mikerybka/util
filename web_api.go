package util

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// The WebAPI type represents an WebAPI backed by any JSON-serializable
// Go object.
type WebAPI struct {
	Types    map[string]Type
	RootType string
	Data     FileSystem
}

func (api *WebAPI) FS() *StructuredFilesystem {
	return &StructuredFilesystem{
		Types:    api.Types,
		RootType: api.RootType,
		Data:     api.Data,
	}
}

// func (api *WebAPI) Type() Type {
// 	if strings.HasPrefix(api.RootType, "map[string]") {
// 		return Type{
// 			IsMap: true,
// 			ElemType: ,
// 		}
// 	}

// 	if strings.HasPrefix(api.RootType, "*") {

// 	}

// 	if strings.HasPrefix(api.RootType, "[]") {

// 	}

// 	t, ok := api.Types[api.RootType]
// 	if !ok {
// 		panic(fmt.Sprintf("invalid type: %s", api.RootType))
// 	}
// 	return t
// }

// ServeHTTP serves a generic REST API.
func (api *WebAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If the request is to the special /meta endpoint, we return the type info.
	if r.URL.Path == "/meta" {
		json.NewEncoder(w).Encode(Metadata{
			Type:  api.RootType,
			Types: api.Types,
		})
		return
	}

	api.FS().ServeHTTP(w, r)

	// // Parse the path into sections.
	// path := ParsePath(r.URL.Path)

	// // If we're at the root, we serve it.
	// if len(path) == 0 {
	// 	api.ServeRoot(w, r)
	// 	return
	// }

	// // If not, we dig one level deeper.
	// first := path[0]
	// rest := "/" + strings.Join(path[1:], "/")

	// // If the type doesn't have nesting, return not found.
	// if api.Type().IsScalar {
	// 	api.ServeNotFound(w, r)
	// 	return
	// }

	// // If it's a pointer, same thing.
	// if api.Type().IsPointer {
	// 	api.ServeNotFound(w, r)
	// 	return
	// }

	// // If the data is an array
	// if api.Type().IsArray {
	// 	// As a special case, handle POST requests by appending to the array.
	// 	if rest == "/" && r.Method == "POST" {
	// 		// Add to the end of the array
	// 		api.ServePOST(w, r)
	// 		return
	// 	}

	// 	// Otherwise, dig one level deeper.
	// 	next := &WebAPI{
	// 		Types:    api.Types,
	// 		RootType: api.Type().ElemType,
	// 		Data:     api.Data.Dig(first),
	// 	}
	// 	r.URL.Path = rest
	// 	next.ServeHTTP(w, r)
	// 	return
	// }

	// // If the data is a map, dig.
	// if api.Type().IsMap {
	// 	next := &WebAPI{
	// 		Types:    api.Types,
	// 		RootType: api.Type().ElemType,
	// 		Data:     api.Data.Dig(first),
	// 	}
	// 	r.URL.Path = rest
	// 	next.ServeHTTP(w, r)
	// 	return
	// }

	// if api.Type().IsStruct {
	// 	for _, f := range api.Type().Fields {
	// 		if f.ID == first {
	// 			next := &WebAPI{
	// 				Types:    api.Types,
	// 				RootType: api.Type().ElemType,
	// 				Data:     api.Data.Dig(first),
	// 			}
	// 			r.URL.Path = rest
	// 			next.ServeHTTP(w, r)
	// 			return
	// 		}
	// 	}

	// 	// If the field is not found, return not found.
	// 	api.ServeNotFound(w, r)
	// }

	// panic("invalid type")
}

func (api *WebAPI) ServeNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("null"))
}

func (api *WebAPI) ServePOST(w http.ResponseWriter, r *http.Request) {
	// Get the size of the dir
	entries, err := api.Data.ReadDir("/")
	if err != nil {
		panic(err)
	}
	dirSize := len(entries)

	// Figure out the next ID
	nextID := strconv.Itoa(dirSize)

	// Read the data from the request body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Create a new object with that ID
	err = api.Data.WriteFile(nextID, b)
	if err != nil {
		panic(err)
	}
}

// func (api *WebAPI) ServeRoot(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		// If the value is a scalar or a pointer, read from file.
// 		if api.Type().IsScalar || api.Type().IsPointer {
// 			b, err := api.Data.ReadFile("/")
// 			if err != nil {
// 				panic(err)
// 			}
// 			w.Write(b)
// 			return
// 		}

// 		// If the value is an array or a map, read the directory.
// 		if api.Type().IsArray || api.Type().IsMap {
// 			entries, err := api.Data.ReadDir("/")
// 			if err != nil {
// 				panic(err)
// 			}
// 			json.NewEncoder(w).Encode(entries)
// 			return
// 		}

// 		// If the value is a struct, build a JSON object.
// 		if api.Type().IsStruct {
// 			// Open the object.
// 			w.Write([]byte("{"))

// 			// Write each field out.
// 			for i, f := range api.Type().Fields {
// 				// Add a comma if it's not the first field.
// 				if i > 0 {
// 					fmt.Fprintf(w, ",")
// 				}

// 				// Write the field ID as the key.
// 				fmt.Fprintf(w, "\"%s\": ", f.ID)

// 				// Build a new API object.
// 				fieldAPI := &WebAPI{
// 					Types:    api.Types,
// 					RootType: f.Type,
// 					Data:     api.Data.Dig(f.ID),
// 				}

// 				// Serve the field.
// 				fieldAPI.ServeHTTP(w, r)
// 			}

// 			// Close object.
// 			w.Write([]byte("}"))
// 			return
// 		}

// 		panic("invalid type")
// 	}

// 	if r.Method == "PUT" {
// 		// If the value is a scalar or a pointer, write to file.
// 		if api.Type().IsScalar || api.Type().IsPointer {
// 			// TODO
// 			return
// 		}

// 		// Otherwise, don't allow the request.
// 	}

// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// }
