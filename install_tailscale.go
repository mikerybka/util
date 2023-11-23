package util

import (
	"fmt"
	"os/exec"
)

func InstallTailscale() error {
	osid, _ := GetOSID()
	if osid == "fedora" {
		cmd := exec.Command("dnf", "config-manager", "--add-repo", "https://pkgs.tailscale.com/stable/fedora/tailscale.repo")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}

		cmd = exec.Command("dnf", "install", "-y", "tailscale")
		out, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}

		cmd = exec.Command("systemctl", "enable", "--now", "tailscaled")
		out, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}

		return nil
	}

	panic("not implemented")
}
