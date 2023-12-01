package util

import (
	"os"
	"runtime"
	"strings"
)

func OS() string {
	if runtime.GOOS == "linux" {
		b, err := os.ReadFile("/etc/os-release")
		if err != nil {
			return "linux"
		}
		for _, line := range strings.Split(string(b), "\n") {
			if strings.HasPrefix(line, "ID=") {
				return strings.TrimPrefix(line, "ID=")
			}
		}
		return "linux"
	}
	return runtime.GOOS
}
