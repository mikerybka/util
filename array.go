package util

import (
	"encoding/json"
	"path/filepath"
	"strconv"
)

type Array[T any] struct {
	Path  []string
	Value []T
}

func (a *Array[T]) ID() string {
	return JoinPath(a.Path)
}

func (a *Array[T]) JSON() string {
	b, err := json.Marshal(a.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (a *Array[T]) Type() string {
	return "array"
}

func (a *Array[T]) Ptr() any {
	return a.Value
}

func (a *Array[T]) Dig(s string) (Object, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, false
	}
	return NewObject(filepath.Join(a.ID(), s), a.Value[i])
}
