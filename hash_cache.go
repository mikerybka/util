package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/pkg/util"
)

type HashCache struct {
	DataDir string
	cache   map[[32]byte][]byte
}

func (s *HashCache) Read(hash [32]byte) ([]byte, error) {
	b, ok := s.cache[hash]
	if ok {
		return b, nil
	}
	b, err := os.ReadFile(s.path(hash))
	if err != nil {
		return nil, err
	}
	s.cache[hash] = b
	return b, nil
}

func (s *HashCache) Exists(hash [32]byte) (bool, error) {
	_, ok := s.cache[hash]
	if ok {
		return true, nil
	}
	_, err := os.Stat(s.path(hash))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *HashCache) path(hash [32]byte) string {
	return filepath.Join(s.DataDir, hex.EncodeToString(hash[:]))
}

func (s *HashCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, _, isRoot := util.PopPath(r.URL.Path)

	if r.Method == "GET" {
		if isRoot {
			http.NotFound(w, r)
			return
		}
		hashBytes, err := hex.DecodeString(first)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		var hash [32]byte
		copy(hash[:], hashBytes)
		b, err := s.Read(hash)
		if errors.Is(err, os.ErrNotExist) {
			http.NotFound(w, r)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
		return
	}

	if r.Method == "POST" {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hash := sha256.Sum256(b)
		exists, err := s.Exists(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if exists {
			return
		}
		path := s.path(hash)
		err = util.WriteFile(path, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, hex.EncodeToString(hash[:]))
		return
	}
}
