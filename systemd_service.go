package util

import (
	"fmt"
	"os/exec"
)

type SystemdService struct {
	Name string
}

func (s *SystemdService) Restart() error {
	cmd := exec.Command("systemctl", "restart", s.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return fmt.Errorf("%s: %s", err, out)
		}
		return err
	}
	if len(out) > 0 {
		return fmt.Errorf("systemctl restart %s: %s", s.Name, out)
	}
	return nil
}
