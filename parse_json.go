package util

import "encoding/json"

func ParseJSON[T any](s string) T {
	var v T
	json.Unmarshal([]byte(s), &v)
	return v
}
