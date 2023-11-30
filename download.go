package util

import (
	"io"
	"net/http"
	"os"
)

func Download(url, target string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
