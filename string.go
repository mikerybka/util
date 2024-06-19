package util

import (
	"encoding/json"
)

type String struct {
	Path  []string
	Value string
}

func (s *String) ID() string {
	return JoinPath(s.Path)
}

func (s *String) JSON() string {
	b, err := json.Marshal(s.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *String) Type() string {
	return "string"
}

func (s *String) Ptr() any {
	return &s.Value
}

func (s *String) Dig(p string) (Object, bool) {
	return nil, false
}
