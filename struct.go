package util

import (
	"encoding/json"
)

type Struct struct {
	Path  []string
	Value map[string]any
}

func (s *Struct) ID() string {
	return JoinPath(s.Path)
}

func (s *Struct) JSON() string {
	b, err := json.Marshal(s.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *Struct) Type() string {
	return "struct"
}

func (s *Struct) Ptr() any {
	return s.Value
}

func (s *Struct) Dig(p string) (Object, bool) {
	v, ok := s.Value[p]
	return NewObject(append(s.Path, p), v), ok
}
