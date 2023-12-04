package util

import (
	"encoding/json"
	"io/fs"
	"os"
)

func ReadJSONFile(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}

func FReadJSONFile(fsys fs.FS, path string, v any) error {
	f, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}
