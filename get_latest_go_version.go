package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetLatestGoVersion() (string, error) {
	res, err := http.Get("https://go.dev/dl/?mode=json")
	if err != nil {
		return "", err
	}
	d := []struct {
		Version string `json:"version"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return "", err
	}
	if len(d) < 1 {
		return "", fmt.Errorf("no versions listed")
	}
	return strings.TrimPrefix(d[0].Version, "go"), nil
}
