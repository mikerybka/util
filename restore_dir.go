package util

import (
	"os"
)

func RestoreDir(backupFile string, encKey string, outDir string) error {
	f, err := os.Open(backupFile)
	if err != nil {
		return err
	}
	defer f.Close()
	err = DecryptAndDecompressDir(f, outDir, encKey)
	if err != nil {
		return err
	}
	return nil
}
