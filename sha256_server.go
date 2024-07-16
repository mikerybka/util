package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type SHA256Server struct {
	FileDir string
	User    string
	Pass    string
}

func (s *SHA256Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok {
		http.NotFound(w, r)
		return
	}
	if user != s.User || pass != s.Pass {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		// Don't allow index requests.
		if len(ParsePath(r.URL.Path)) != 1 {
			http.NotFound(w, r)
			return
		}

		http.FileServer(http.Dir(s.FileDir)).ServeHTTP(w, r)
	}

	if r.Method == http.MethodPost {
		// Create a SHA-256 hash object
		hash := sha256.New()

		// Create a temporary file
		tmpPath := filepath.Join(s.FileDir, UnixNanoTimestamp())
		err := os.MkdirAll(filepath.Dir(tmpPath), os.ModePerm)
		if err != nil {
			panic(err)
		}
		tmpFile, err := os.Create(tmpPath)
		if err != nil {
			panic(err)
		}
		defer os.Remove(tmpPath)

		// Create a buffer to read chunks
		buf := make([]byte, 1024*1024) // 1MB buffer
		for {
			// Read a chunk from the request body
			n, err := r.Body.Read(buf)
			if n > 0 {
				// Write the chunk to the hash
				if _, err := hash.Write(buf[:n]); err != nil {
					panic(err)
				}
				// Write the chunk to the temporary file
				if _, err := tmpFile.Write(buf[:n]); err != nil {
					panic(err)
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}

		// Compute the SHA-256 hash
		hashSum := hash.Sum(nil)
		hashHex := hex.EncodeToString(hashSum)

		// Close the temporary file to flush all writes to disk
		if err := tmpFile.Close(); err != nil {
			panic(err)
		}

		// Rename the temporary file to the SHA-256 hash name
		finalFileName := filepath.Join(s.FileDir, hashHex)
		if err := os.Rename(tmpPath, finalFileName); err != nil {
			panic(err)
		}

		// Return the hash
		fmt.Fprintln(w, hashHex)
	}
}
