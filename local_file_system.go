package util

import (
	"errors"
	"os"
	"path/filepath"
)

type LocalFileSystem struct {
	Root string
}

func (fs *LocalFileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(filepath.Join(fs.Root, path))
}

func (fs *LocalFileSystem) ReadDir(path string) ([]string, error) {
	res := []string{}
	entries, err := os.ReadDir(filepath.Join(fs.Root, path))
	if errors.Is(err, os.ErrNotExist) {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		res = append(res, entry.Name())
	}
	return res, nil
}

func (fs *LocalFileSystem) IsDir(path string) bool {
	fi, _ := os.Stat(filepath.Join(fs.Root, path))
	return fi.IsDir()
}

func (fs *LocalFileSystem) IsFile(path string) bool {
	fi, err := os.Stat(filepath.Join(fs.Root, path))
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	if err != nil {
		panic(err)
	}
	return !fi.IsDir()
}

func (fs *LocalFileSystem) WriteFile(path string, b []byte) error {
	err := os.MkdirAll(filepath.Dir(filepath.Join(fs.Root, path)), os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(fs.Root, path), b, os.ModePerm)
}
func (fs *LocalFileSystem) MakeDir(path string) error {
	return os.MkdirAll(filepath.Join(fs.Root, path), os.ModePerm)
}
func (fs *LocalFileSystem) Remove(path string) error {
	return os.RemoveAll(filepath.Join(fs.Root, path))
}
