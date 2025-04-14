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
	  <meta charset="UTF-8" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	  <title>Password Protected</title>
	  <script src="https://cdn.tailwindcss.com"></script>
	</head>
	<body class="bg-gray-100 min-h-screen flex items-center justify-center">
	  <div id="form-wrapper" class="w-full max-w-sm p-6 bg-white rounded-2xl shadow-lg transition-transform duration-300 transform">
		<form method="POST" action="/auth/login?target=%s">
		  <h1 class="text-lg font-semibold text-gray-800 mb-4">This site is password protected.</h1>
		  <input
			id="password"
			name="password"
			type="password"
			placeholder="Enter password"
			class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black mb-4"
			required
		  />
		  <button type="submit" class="w-full bg-black text-white py-2 rounded-lg hover:bg-gray-900 transition">Enter</button>
		</form>
	  </div>
	
	  <script>
		const input = document.getElementById('password');
		const wrapper = document.getElementById('form-wrapper');
	
		input.addEventListener('focus', () => {
		  wrapper.classList.add('-translate-y-24');
		});
	
		input.addEventListener('blur', () => {
		  wrapper.classList.remove('-translate-y-24');
		});
	  </script>
	</body>
	</html>
	
`, r.URL)))
}
