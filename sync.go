package util

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

func Sync(src, dst string) error {
	src = strings.TrimSuffix(src, "/") + "/"

	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("rsync", "-avnh", "--delete", src, dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	cmd = exec.Command("rsync", "-a", "--delete", src, dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
