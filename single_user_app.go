package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

type SingleUserApp[T http.Handler] struct {
	Twilio    *TwilioClient
	UserPhone string
	Files     FileSystem
}

func (app *SingleUserApp[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dataLocation := "app/data.json"
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

	b, _ := app.Files.ReadFile(dataLocation)
	var data T
	json.Unmarshal(b, data)

	data.ServeHTTP(w, r)

	if IsMutation(r) {
		b, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			panic(err)
		}
		err = app.Files.WriteFile(dataLocation, b)
		if err != nil {
			panic(err)
		}
	}
}

func (a *SingleUserApp[T]) AuthHandler() *SingleUserAuthHandler {
	return &SingleUserAuthHandler{
		UserPhone: a.UserPhone,
		AuthFiles: a.Files.Dig("auth"),
		Twilio:    a.Twilio,
	}
}

func (a *SingleUserApp[T]) AuthFiles() FileSystem {
	return a.Files.Dig("auth")
}

func (a *SingleUserApp[T]) Authorized(r *http.Request) (bool, error) {
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

func (a *SingleUserApp[T]) GetSession(token string) (*Session, error) {
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
