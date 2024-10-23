package util

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Word string

func (w Word) String() string {
	return string(w)
}

func (w Word) Title() Word {
	l := language.English
	return Word(cases.Title(l).String(w.String()))
}

func (w Word) Lower() Word {
	return Word(strings.ToLower(w.String()))
}

func (w Word) Upper() Word {
	return Word(strings.ToUpper(w.String()))
}

func (w Word) StripNonAlphaNumeric() Word {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	s := ""
	for _, ch := range string(w) {
		if strings.ContainsRune(charset, ch) {
			s += string(ch)
		}
	}
	return Word(s)
}
