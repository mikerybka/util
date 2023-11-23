package util

import (
	"fmt"
	"os/exec"
)

func AuthenticateTailscale(authkey string) error {
	cmd := exec.Command("tailscale", "up", "--authkey", authkey)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}
