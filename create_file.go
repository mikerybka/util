package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(path string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("mkdirall: %s", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create: %s", err)
	}
	return f, nil
}
