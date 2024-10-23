package util

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email        string
	PasswordHash string
}

func (u *User) IsPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(pass))
	return err == nil
}

func (u *User) SetPassword(pass string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.PasswordHash = string(hashedPassword)
}
