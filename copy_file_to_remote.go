package util

import (
	"fmt"
	"os/exec"
)

func CopyFileToRemote(user, addr, src, dst string) ([]byte, error) {
	c := exec.Command(
		"scp",
		"-o", "StrictHostKeyChecking=no",
		src,
		fmt.Sprintf("%s@%s:%s", user, addr, dst),
	)
	return c.CombinedOutput()
}
