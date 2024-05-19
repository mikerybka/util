package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

type MultiUserApp struct {
	Twilio    *TwilioClient
	AuthFiles FileSystem
	App       http.Handler
}

func (a *MultiUserApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/auth") {
		http.StripPrefix("/auth", a.AuthHandler()).ServeHTTP(w, r)
		return
	}

	allowed, userID, err := a.Authorized(r)
	if err != nil {
		panic(err)
	}
	if !allowed {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	r.Header.Add("User", userID)
	a.App.ServeHTTP(w, r)
}

func (a *MultiUserApp) AuthHandler() *AuthHandler {
	return &AuthHandler{
		AuthFiles: a.AuthFiles,
		Twilio:    a.Twilio,
	}
}

func (a *MultiUserApp) Authorized(r *http.Request) (bool, string, error) {
	token := r.Header.Get("Token")
	session, err := a.GetSession(token)
	if errors.Is(err, os.ErrNotExist) {
		return false, "", nil
	}
	if err != nil {
		return false, "", err
	}

	path := ParsePath(r.URL.Path)
	if len(path) == 0 {
		if r.Method == "POST" ||
			r.Method == "PUT" ||
			r.Method == "PATCH" ||
			r.Method == "DELETE" {
			return false, session.UserID, nil
		} else {
			return true, session.UserID, nil
		}
	}

	orgID := path[0]
	org, err := a.GetOrg(orgID)
	if err != nil {
		return false, session.UserID, err
	}

	return org.HasMember(session.UserID), session.UserID, nil
}

func (a *MultiUserApp) GetSession(token string) (*Session, error) {
	b, err := a.AuthFiles.ReadFile("/sessions/" + token)
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

func (a *MultiUserApp) GetOrg(id string) (*Org, error) {
	b, err := a.AuthFiles.ReadFile("/orgs/" + id)
	if err != nil {
		return nil, err
	}
	o := &Org{}
	err = json.Unmarshal(b, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}
