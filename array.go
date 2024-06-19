package util

import (
	"encoding/json"
	"strconv"
)

type Array struct {
	Path  []string
	Value []any
}

func (a *Array) ID() string {
	return JoinPath(a.Path)
}

func (a *Array) JSON() string {
	b, err := json.Marshal(a.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (a *Array) Type() string {
	return "array"
}

func (a *Array) Ptr() any {
	return a.Value
}

func (a *Array) Dig(s string) (Object, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, false
	}
	return NewObject(append(a.Path, s), a.Value[i]), true
}
