package util

import (
	"encoding/json"
	"io/fs"
	"os"
)

func ReadJSONFile[T any](path string) T {
	var v T
	b, _ := os.ReadFile(path)
	json.Unmarshal(b, v)
	return v
}

func FReadJSONFile(fsys fs.FS, path string, v any) error {
	f, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}
