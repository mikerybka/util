package util

import (
	"os"
	"path/filepath"
)

func CreateP(name string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(name), os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
