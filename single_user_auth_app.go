package util

import (
	"net/http"
	"strings"
)

type SingleUserAuthApp struct {
	Files FileSystem
	App   http.Handler
}

func (a *SingleUserAuthApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/auth") {
		http.StripPrefix("/auth", a.AuthHandler()).ServeHTTP(w, r)
		return
	}

	user, pass, ok := r.BasicAuth()
	if !ok || !a.Authorized(user, pass) {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	a.App.ServeHTTP(w, r)
}

func (a *SingleUserAuthApp) Authorized(token string) bool {
	b, err := a.Files.ReadFile("/admin/id")
	if err != nil {
		panic(err)
	}
	adminID := string(b)
	if userID != adminID {
		return false
	}
	return a.Files.IsFile("/admin/tokens/" + pass)
}

type SingleUserAuth struct {
	Files FileSystem
}

func (h *SingleUserAuth) SendLoginCode(phone string) error
func (h *SingleUserAuth) Login(code string) (string, error)
func (h *SingleUserAuth) Logout(token string) error
func (h *SingleUserAuth) WhoAmI(token string) (string, error)
