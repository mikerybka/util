package util

import "encoding/json"

type Session struct {
	UserID string
	Token  string
}

func (s *Session) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}
