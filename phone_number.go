package util

import (
	"fmt"
	"unicode"
)

type PhoneNumber string

func (phone PhoneNumber) Validate() error {
	for _, ch := range phone {
		if !unicode.IsDigit(ch) {
			return fmt.Errorf("all chars must be digits")
		}
	}
	return nil
}
