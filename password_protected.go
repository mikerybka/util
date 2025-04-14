package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func PasswordProtected(passwd string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle logins
		if r.Method == "POST" && r.URL.Path == "/auth/login" {
			// Get inputs
			target := r.URL.Query().Get("target")
			if target == "" {
				target = "/"
			}
			req := &struct {
				Password string `json:"password"`
			}{}
			if util.ContentType(r, "application/json") {
				err := json.NewDecoder(r.Body).Decode(req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			} else {
				req.Password = r.FormValue("password")
			}

			// Check password
			if req.Password != passwd {
				http.Error(w, "wrong password", http.StatusBadRequest)
				return
			}

			// Set cookie
			http.SetCookie(w, &http.Cookie{
				Name:  "password",
				Value: req.Password,
				Path:  "/",
			})

			// Redirect
			http.Redirect(w, r, target, http.StatusSeeOther)
			return
		}

		// Get password from either the header or a cookie
		pass := r.Header.Get("Password")
		if pass == "" {
			c, err := r.Cookie("password")
			if err != nil {
				passwordInputPage(w, r)
				return
			}
			pass = c.Value
		}

		// Check password
		if pass != passwd {
			http.SetCookie(w, &http.Cookie{
				Name:  "password",
				Value: "",
			})
			http.NotFound(w, r)
		}

		// Handle request
		h.ServeHTTP(w, r)
	})
}

func passwordInputPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		<head>
		  <meta charset="UTF-8">
		  <title>Password</title>
		</head>
		<body>
		  <h1>Password please.</h1>
		  <form method="POST" action="/auth/login?target=%s">
			<input type="password" id="password" name="password" required>
			<button type="submit">Enter</button>
		  </form>
		</body>
		</html>
		`, r.URL)))
}
