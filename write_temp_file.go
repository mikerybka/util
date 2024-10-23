package util

import "os"

func WriteTempFile(b []byte) (string, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return "", err
	}
	_, err = f.Write(b)
	f.Close()
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
