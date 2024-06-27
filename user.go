package util

import "strings"

type User struct {
	// Public ID
	ID string

	// Contact
	Phone PhoneNumber
	Email Email

	// Personal
	FullName string
}

func (u *User) FirstName() string {
	if u.FullName == "" {
		return ""
	}
	return strings.Split(u.FullName, " ")[0]
}
