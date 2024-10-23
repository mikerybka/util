package util

import "strings"

func NewName(s string) Name {
	words := strings.Split(s, " ")
	name := Name{}
	for _, w := range words {
		if w != "" {
			name = append(name, Word(w))
		}
	}
	return name
}

type Name []Word

func (n Name) String() string {
	s := ""
	for i, w := range n {
		if i > 0 {
			s += " "
		}
		s += w.String()
	}
	return s
}

// ID returns the id friendly string.
// Ex: "Green Button" => "green_button"
func (n Name) ID() string {
	s := ""
	for i, w := range n {
		if i > 0 {
			s += "_"
		}
		s += w.StripNonAlphaNumeric().Lower().String()
	}
	return s
}

// GoExported returns an exported Go name.
// Ex: "Green Button" => "GreenButton"
func (n Name) GoExported() string {
	s := ""
	for _, w := range n {
		s += w.StripNonAlphaNumeric().Title().String()
	}
	return s
}
