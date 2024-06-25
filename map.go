package util

import (
	"encoding/json"
)

type Map struct {
	Path  []string
	Value map[string]any
}

func (m *Map) ID() string {
	return JoinPath(m.Path)
}

func (m *Map) JSON() string {
	b, err := json.Marshal(m.Value)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (m *Map) Type() string {
	return "map"
}

func (m *Map) Ptr() any {
	return m.Value
}

func (m *Map) Dig(s string) (Object, bool) {
	v, ok := m.Value[s]
	if !ok {
		return nil, false
	}
	return NewObject(append(m.Path, s), v), true
}
