package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FileServer struct {
	DB *PostgresDB
}

func (fs *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if r.Method == http.MethodGet {
		data, err := fs.get(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
		return
	}

	if r.Method == http.MethodPut {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fs.put(path, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	http.NotFound(w, r)
}

func (fs *FileServer) get(path string) ([]byte, error) {
	q := `SELECT data FROM files WHERE path = $1`
	data := []byte{}
	err := fs.DB.Open().QueryRow(q, path).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fs *FileServer) put(path string, data []byte) error {
	q := `
		INSERT INTO files (path, data) 
		VALUES ($1, $2)
		ON CONFLICT (path) 
		DO UPDATE SET data = EXCLUDED.data
	`
	res, err := fs.DB.Open().Exec(q, path, data)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row to be updated, %d updated", rowsAffected)
	}
	return nil
}
