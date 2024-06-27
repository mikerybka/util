package util

import "net/mail"

type Email string

func (email Email) Validate() error {
	_, err := mail.ParseAddress(string(email))
	return err
}
