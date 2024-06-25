package util

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type Array struct {
	Path  []string
	Value any
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
	path := append(a.Path, s)
	v := reflect.ValueOf(a.Value)
	if i >= v.Len() || i < 0 {
		return nil, false
	}
	return NewObject(path, v.Index(i).Interface()), true
}
