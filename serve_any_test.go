package util_test

import (
	"testing"
)

type TestType struct {
	A string
	B int
	C bool
	D map[string]string
	E map[string]TestType
	F []string
	G []TestType
}

func GetRootTest(t *testing.T) {

}
