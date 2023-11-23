package util

import (
	"fmt"
	"os"
	"strings"
)

func GetOSID() (string, error) {
	b, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ID=") {
			return strings.TrimPrefix(line, "ID="), nil
		}
	}
	return "", fmt.Errorf("not found")
}
