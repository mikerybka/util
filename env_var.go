package util

import "os"

func EnvVar(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}
