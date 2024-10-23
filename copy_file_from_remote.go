package util

import (
	"fmt"
	"os/exec"
)

func CopyFileFromRemote(user, addr, src, dst string) ([]byte, error) {
	c := exec.Command(
		"scp",
		"-o", "StrictHostKeyChecking=no",
		fmt.Sprintf("%s@%s:%s", user, addr, src),
		dst,
	)
	return c.CombinedOutput()
}
