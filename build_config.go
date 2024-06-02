package util

import "fmt"

type BuildConfig struct {
	Type string
}

func (config *BuildConfig) Build(path string) error {
	switch config.Type {
	default:
		panic(fmt.Sprintf("unknown type: %s", config.Type))
	}
}
