package util

import (
	"net/http"
)

type User struct {
	// Public ID
	ID string

	// Contact
	Phone PhoneNumber
	Email Email

	// Personal
	FirstName string
	LastName  string

	// Data
	Schemas Table[*Schema]
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == "schema.cafe" {
		u.Schemas.ServeHTTP(w, r)
	}
}
