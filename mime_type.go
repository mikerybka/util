package util

import (
	"net/http"
	"os"
)

func MIMEType(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		panic(err)
	}
	return http.DetectContentType(buf)
}
