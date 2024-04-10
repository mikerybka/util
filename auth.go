package util

import (
	"fmt"
	"time"
)

type Auth struct {
	Accounts      map[string]Account
	SessionTokens map[string]string
	LoginCodes    map[string]string
}

func (a *Auth) SendLoginCode(phone string) error {
	if len(phone) != 10 {
		return fmt.Errorf("expected 10 digit phone number")
	}
	acc, ok := a.Accounts[phone]
	if !ok {
		acc = Account{
			Phone: phone,
		}
		a.Accounts[phone] = acc
	}
	code := RandomCode(6)
	a.LoginCodes[code] = phone
	go func() {
		time.Sleep(5 * time.Minute)
		delete(a.LoginCodes, code)
	}()
	return nil
}

func (a *Auth) Login(code string) (string, error) {
	panic("not implemented")
}
func (a *Auth) Logout(token string) error {
	panic("not implemented")
}
func (a *Auth) WhoAmI(token string) (string, error) {
	panic("not implemented")
}
