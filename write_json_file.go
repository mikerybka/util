package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteJSONFile(path string, v any) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, os.ModePerm)
}
