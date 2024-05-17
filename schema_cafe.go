package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type SchemaCafe struct {
	Data FileSystem
}

func (s *SchemaCafe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := ParsePath(r.URL.Path)

	if len(path) == 0 {
		s.Homepage(w, r)
		return
	}

	orgID := path[0]
	org := s.GetOrg(orgID)
	http.StripPrefix("/"+orgID, org).ServeHTTP(w, r)
	if IsMutation(r) {
		err := s.SaveOrg(org)
		if err != nil {
			panic(err)
		}
	}
}

func (s *SchemaCafe) Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "schema.cafe")
}

func (s *SchemaCafe) SaveOrg(org *SchemaCafeOrg) error {
	path := "/" + org.ID
	b, err := json.Marshal(org)
	if err != nil {
		panic(err)
	}
	return s.Data.WriteFile(path, b)
}

func (s *SchemaCafe) GetOrg(id string) *SchemaCafeOrg {
	// Read /:id
	b, err := s.Data.ReadFile("/" + id)

	// Return nil if file doesn't exist.
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	// Panic on any other error.
	if err != nil {
		panic(err)
	}

	// Read the data into an object.
	org := &SchemaCafeOrg{}
	err = json.Unmarshal(b, org)

	// This should never happen.
	if err != nil {
		panic(err)
	}

	// Return it.
	return org
}
