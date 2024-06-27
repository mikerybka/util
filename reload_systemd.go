package util

import (
	"fmt"
	"os/exec"
)

func ReloadSystemd() error {
	cmd := exec.Command("systemctl", "daemon-relooad")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return fmt.Errorf("%s: %s", err, out)
		}
		return err
	}
	if len(out) > 0 {
		return fmt.Errorf("systemctl daemon-relooad: %s", out)
	}
	return nil
}
