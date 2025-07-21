package util

type File[T any] struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Value T      `json:"value"`
}
