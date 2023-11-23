package util

import (
	"fmt"
	"os/exec"
)

func InstallNode() error {
	osid, _ := GetOSID()
	if osid == "fedora" {
		cmd := exec.Command("dnf", "install", "-y", "nodejs")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}
		return nil
	}

	panic("not implemented")
}
