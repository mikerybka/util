package util

import (
	"encoding/json"
)

type Bool struct {
	Path  []string
	Value bool
}

func (b *Bool) ID() string {
	return JoinPath(b.Path)
}

func (b *Bool) JSON() string {
	byt, err := json.Marshal(b.Value)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

func (b *Bool) Type() string {
	return "bool"
}
func (b *Bool) Ptr() any {
	return &b.Value
}
func (b *Bool) Dig(s string) (Object, bool) {
	return nil, false
}
