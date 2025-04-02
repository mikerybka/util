package util

import (
	"os/exec"
	"strings"
)

func GetInstalledGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Split(strings.TrimPrefix(string(out), "go version go"), " ")[0], nil
}
