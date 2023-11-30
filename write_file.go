package util

import (
	"os"
	"path/filepath"
)

func WriteFile(path string, b []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, os.ModePerm)
}
