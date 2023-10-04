package util

import "os"

func PrintJSON(v any) (int, error) {
	return FprintJSON(os.Stdout, v)
}
