package util

import (
	"net/http"
	"time"
)

func DeleteCookie(w http.ResponseWriter, name string) {
	// Set the cookie with the same name and an expiration time in the past
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})
}
