package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

type SingleUserApp struct {
	Twilio    *TwilioClient
	UserPhone string
	Files     FileSystem
	Handler   http.Handler
}

func (app *SingleUserApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/auth") {
		http.StripPrefix("/auth", app.AuthHandler()).ServeHTTP(w, r)
		return
	}

	allowed, err := app.Authorized(r)
	if err != nil {
		panic(err)
	}
	if !allowed {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	app.Handler.ServeHTTP(w, r)
}

func (a *SingleUserApp) AuthHandler() *SingleUserAuthHandler {
	return &SingleUserAuthHandler{
		UserPhone: a.UserPhone,
		AuthFiles: a.Files.Dig("auth"),
		Twilio:    a.Twilio,
	}
}

func (a *SingleUserApp) AuthFiles() FileSystem {
	return a.Files.Dig("auth")
}

func (a *SingleUserApp) Authorized(r *http.Request) (bool, error) {
	token := r.Header.Get("Token")
	_, err := a.GetSession(token)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *SingleUserApp) GetSession(token string) (*Session, error) {
	b, err := a.AuthFiles().ReadFile("/sessions/" + token)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	s := &Session{}
	err = json.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
