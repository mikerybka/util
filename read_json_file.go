package util

import (
	"encoding/json"
	"io/fs"
	"os"
)

func ReadJSONFile(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func FReadJSONFile(fsys fs.FS, path string, v any) error {
	f, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}
