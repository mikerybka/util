package util

import (
	"io"
)

type ReactElement struct {
	Type  string
	Props []struct {
		Key   string
		Value any
	}
	Children []ReactElement
}

func (el *ReactElement) WriteNextJSPage(w io.Writer) error {
	s := `import React from 'react';
import { ` + el.Type + ` } from '@brass.software/components';

export default function Page() {
	return <` + el.Type + ` />
}`
	_, err := w.Write([]byte(s))
	return err
}
