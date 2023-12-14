package util

import (
	"os"
	"path/filepath"
)

func BackupDir(path string, encKey string, outfile string) error {
	err := os.MkdirAll(filepath.Dir(outfile), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer f.Close()
	err = EncryptAndCompressDir(path, f, encKey)
	if err != nil {
		return err
	}
	return nil
}
