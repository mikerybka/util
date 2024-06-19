package util

import (
	"encoding/json"
)

type Int struct {
	Path  []string
	Value int64
}

func (i *Int) ID() string {
	return JoinPath(i.Path)
}

func (i *Int) JSON() string {
	b, err := json.Marshal(i.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}
func (i *Int) Type() string {
	return "int"
}
func (i *Int) Ptr() any {
	return &i.Value
}
func (i *Int) Dig(s string) (Object, bool) {
	return nil, false
}
