package util

import (
	"net/http"
	"strings"
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
	Schemas *Table[*Schema]
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Redirect if unauthenticated
	if u == nil {
		http.Redirect(w, r, "/auth/send-login-code", http.StatusFound)
		return
	}

	// special case for dev
	host := r.Host
	if strings.HasPrefix(host, "localhost") {
		first, rest, isRoot := PopPath(r.URL.Path)
		if isRoot {
			http.NotFound(w, r)
			return
		}
		r.URL.Path = rest
		host = first
	}

	// Serve each app in the system
	switch host {
	case "schema.cafe", "www.schema.cafe":
		u.SchemaCafe(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (u *User) SchemaCafe(w http.ResponseWriter, r *http.Request) {
	u.Schemas.ServeHTTP(w, r)
}
