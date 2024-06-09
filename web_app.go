package util

import "net/http"

type WebApp struct {
	Name             string
	Description      string
	Author           string
	Keywords         []string
	Favicon          []byte
	Icon             []byte
	Types            map[string]Type
	CoreResourceType string
	TwilioClient     *TwilioClient
	Files            FileSystem
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := ParsePath(r.URL.Path)
	if len(p) == 0 {
		switch r.Method {
		case "GET":
			app.GetRoot(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}

	if p[0] == "auth" && len(p) == 2 {
		switch p[1] {
		case "login":
			switch r.Method {
			case "GET":
				app.GetAuthLogin(w, r)
			case "POST":
				app.PostAuthLogin(w, r)
			default:
				http.NotFound(w, r)
			}
		case "send-login-code":
			switch r.Method {
			case "GET":
				app.GetAuthSendLoginCode(w, r)
			case "POST":
				app.PostAuthSendLoginCode(w, r)
			default:
				http.NotFound(w, r)
			}
		}
		return
	}

	if p[0] == "favicon.ico" && len(p) == 1 && r.Method == "GET" {
		app.GetFavicon(w, r)
		return
	}

	if p[0] == "meta" && len(p) == 1 && r.Method == "GET" {
		app.GetMeta(w, r)
		return
	}

	if len(p) == 1 {
		switch r.Method {
		case "GET":
			app.GetOrg(w, r)
		case "PUT":
			app.PutOrg(w, r)
		case "POST":
			app.PostOrg(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}

	switch r.Method {
	case "GET":
		app.GetPath(w, r)
	case "PUT":
		app.PutPath(w, r)
	case "PATCH":
		app.PatchPath(w, r)
	case "POST":
		app.PostPath(w, r)
	case "DELETE":
		app.DeletePath(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (app *WebApp) GetRoot(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PostRoot(w http.ResponseWriter, r *http.Request)
func (app *WebApp) GetOrg(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PutOrg(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PostOrg(w http.ResponseWriter, r *http.Request)
func (app *WebApp) GetPath(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PutPath(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PostPath(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PatchPath(w http.ResponseWriter, r *http.Request)
func (app *WebApp) DeletePath(w http.ResponseWriter, r *http.Request)
func (app *WebApp) GetMeta(w http.ResponseWriter, r *http.Request)
func (app *WebApp) GetFavicon(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PostAuthLogin(w http.ResponseWriter, r *http.Request)
func (app *WebApp) GetAuthLogin(w http.ResponseWriter, r *http.Request)
func (app *WebApp) PostAuthSendLoginCode(w http.ResponseWriter, r *http.Request)
func (app *WebApp) GetAuthSendLoginCode(w http.ResponseWriter, r *http.Request)
