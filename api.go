package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func NewAPI[T any](workdir string) *API[T] {
	return &API[T]{
		UserDir:    filepath.Join(workdir, "users"),
		SessionDir: filepath.Join(workdir, "sessions"),
		OrgDir:     filepath.Join(workdir, "orgs"),
		DataDir:    filepath.Join(workdir, "data"),
	}
}

type API[T any] struct {
	UserDir    string
	SessionDir string
	OrgDir     string
	DataDir    string
	locks      map[string]*sync.Mutex
}

func (api *API[T]) Users() *Table[User] {
	return &Table[User]{
		DataDir:  filepath.Join(api.UserDir, "data"),
		IndexDir: filepath.Join(api.UserDir, "indexes"),
		Indexes: Set[string]{
			"Email": true,
		},
	}
}
func (api *API[T]) Orgs() *Table[Org] {
	return &Table[Org]{
		DataDir:  filepath.Join(api.OrgDir, "data"),
		IndexDir: filepath.Join(api.OrgDir, "indexes"),
	}
}
func (api *API[T]) Data() *Table[T] {
	return &Table[T]{
		DataDir:  filepath.Join(api.DataDir, "data"),
		IndexDir: filepath.Join(api.DataDir, "indexes"),
	}
}
func (api *API[T]) Sessions() *Table[string] {
	return &Table[string]{
		DataDir:  filepath.Join(api.SessionDir, "data"),
		IndexDir: filepath.Join(api.SessionDir, "indexes"),
	}
}

func (api *API[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/login", api.Login)
	mux.HandleFunc("POST /auth/logout", api.Logout)
	mux.HandleFunc("POST /auth/register", api.Register)
	mux.HandleFunc("POST /{$}", api.NewOrg)
	mux.HandleFunc("/", api.ServeData)
	mux.ServeHTTP(w, r)
}

func (api *API[T]) Login(w http.ResponseWriter, r *http.Request) {
	// Read request
	req := struct {
		Email    string
		Password string
	}{}
	json.NewDecoder(r.Body).Decode(&req)

	// Find user
	usersWithEmail, err := api.Users().FindIDsBy("Email", req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(usersWithEmail) == 0 {
		http.Error(w, "bad login", http.StatusBadRequest)
		return
	}
	if len(usersWithEmail) > 1 {
		panic("multiple users with the same email")
	}
	userID := ""
	for id := range usersWithEmail {
		userID = id
	}
	user, err := api.Users().Get(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Test password
	if !user.IsPassword(req.Password) {
		http.Error(w, "bad login", http.StatusBadRequest)
		return
	}

	// Create a new session
	sessionID := RandomToken(32)
	err = api.Sessions().Set(sessionID, &userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return session ID
	w.Write([]byte(sessionID))
}

func (api *API[T]) Logout(w http.ResponseWriter, r *http.Request) {
	// Read request
	req := struct {
		SessionID string
	}{}
	json.NewDecoder(r.Body).Decode(&req)

	// Delete session
	err := api.Sessions().Delete(req.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API[T]) Register(w http.ResponseWriter, r *http.Request) {
	// Read request
	req := struct {
		Email    string
		Password string
	}{}
	json.NewDecoder(r.Body).Decode(&req)

	// Check if email is already used
	ids, err := api.Users().FindIDsBy("Email", req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(ids) > 0 {
		http.Error(w, "email already registered", http.StatusBadRequest)
		return
	}

	// Check password length
	if len(req.Password) < 8 {
		http.Error(w, "password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	// Create new user
	id := RandomToken(8)
	for {
		api.Users().Get(id)
		if errors.Is(err, os.ErrNotExist) {
			break
		}
		id = RandomToken(8)
	}
	user := &User{
		Email: req.Email,
	}
	user.SetPassword(req.Password)

	// Save user
	err = api.Users().Set(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API[T]) NewOrg(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (api *API[T]) ServeData(w http.ResponseWriter, r *http.Request) {
	// Get Session
	sessionID := r.Header.Get("SessionID")
	userID, err := api.Sessions().Get(sessionID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get Org
	p := ParsePath(r.URL.Path)
	if len(p) == 0 {
		http.NotFound(w, r)
		return
	}
	orgID := p[0]
	org, err := api.Orgs().Get(orgID)
	if err != nil {
		panic(err)
	}

	// Check authorization
	if IsMutation(r) {
		if !org.IsWriter(*userID) {
			http.NotFound(w, r)
			return
		}
	} else {
		if !org.IsReader(*userID) {
			http.NotFound(w, r)
			return
		}
	}

	// Read data
	d, err := api.Data().Get(orgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Lock if mutation to avoid race conditiions
	if IsMutation(r) {
		api.locks[orgID].Lock()
	}

	// Handle the request
	ServeAny(d, w, r)

	// Save changes if needed
	if IsMutation(r) {
		err := api.Data().Set(orgID, d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		api.locks[orgID].Unlock()
	}
}
