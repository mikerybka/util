package util

import (
	"fmt"
	"os/exec"
)

func Run(cmd *exec.Cmd) error {
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
	}
	return err
}
