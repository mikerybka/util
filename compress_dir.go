package util

import (
	"os"
	"os/exec"
)

// CompressDir compresses the specified directory into a .tar.gz file.
func CompressDir(dir, outFile string) error {
	cmd := exec.Command("tar", "-czvf", outFile, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
