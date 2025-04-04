package util

import "os"

func ReadDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	res := []string{}
	for _, e := range entries {
		res = append(res, e.Name())
	}
	return res, nil
}
