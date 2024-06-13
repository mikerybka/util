package util

import (
	_ "embed"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type WebApp[T any] struct {
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

func (app *WebApp[T]) KeywordString() string {
	return strings.Join(app.Keywords, ",")
}

func (app *WebApp[T]) Metadata() *Metadata {
	return &Metadata{
		Type:  app.CoreResourceType,
		Types: app.Types,
	}
}

func (app *WebApp[T]) AuthHandler() *MultiUserAuthHandler {
	return &MultiUserAuthHandler{
		Twilio:    app.TwilioClient,
		AuthFiles: app.Files.Dig("auth"),
	}
}

func (app *WebApp[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

//go:embed web_templates/home.html
var homeTmpl string

//go:embed web_templates/org.html
var orgTmpl string

//go:embed web_templates/auth/login.html
var loginTmpl string

//go:embed web_templates/auth/send-login-code.html
var sendLoginCodeTmpl string

func (app *WebApp[T]) OrgTmpl() *template.Template {
	return template.Must(template.New("org").Parse(orgTmpl))
}

func (app *WebApp[T]) GetRoot(w http.ResponseWriter, r *http.Request) {
	template.Must(template.New("home").Parse(homeTmpl)).Execute(w, app)
}

func (app *WebApp[T]) PostRoot(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (app *WebApp[T]) GetOrg(w http.ResponseWriter, r *http.Request) {
	p := ParsePath(r.URL.Path)

	orgID := p[0]

	entries, err := app.Files.ReadDir("orgs/" + orgID)
	if errors.Is(err, os.ErrNotExist) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		panic(err)
	}

	template.Must(template.New("org").Parse(orgTmpl)).Execute(w, entries)
}

func (app *WebApp[T]) PutOrg(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (app *WebApp[T]) PostOrg(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (app *WebApp[T]) GetPath(w http.ResponseWriter, r *http.Request) {
	object := Object[T]{}
	path := r.URL.Path
	if app.Files.Dig("schemas").IsDir(path) {
		object.IsDir = true
		entries, err := app.Files.Dig("schemas").ReadDir(path)
		if err != nil {
			panic(err)
		}
		object.Entries = entries
	} else if app.Files.Dig("schemas").IsFile(path) {
		object.IsFile = true
		b, err := app.Files.Dig("schemas").ReadFile(path)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(b, object.Data)
		if err != nil {
			panic(err)
		}
	}
	object.ServeHTTP(w, r)
}

func (app *WebApp[T]) PutPath(w http.ResponseWriter, r *http.Request) {
	var val T
	err := json.NewDecoder(r.Body).Decode(val)
	if err != nil {
		http.Error(w, "bad shape", http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	err = app.Files.Dig("schemas").WriteFile(r.URL.Path, b)
	if err != nil {
		panic(err)
	}
}

func (app *WebApp[T]) PostPath(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (app *WebApp[T]) PatchPath(w http.ResponseWriter, r *http.Request) {
	b, _ := app.Files.Dig("schemas").ReadFile(r.URL.Path)
	var val T
	json.Unmarshal(b, val)
	b, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	err = app.Files.Dig("schemas").WriteFile(r.URL.Path, b)
	if err != nil {
		panic(err)
	}
}
func (app *WebApp[T]) DeletePath(w http.ResponseWriter, r *http.Request) {
	app.Files.Dig("schemas").Remove(r.URL.Path)
}

func (app *WebApp[T]) GetMeta(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(app.Metadata())
}

func (app *WebApp[T]) GetFavicon(w http.ResponseWriter, r *http.Request) {
	w.Write(app.Favicon)
}

func (app *WebApp[T]) PostAuthLogin(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	code := r.FormValue("code")
	session, err := app.AuthHandler().Login(phone, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Token",
		Value: session.Token,
	})
	http.Redirect(w, r, "/"+session.UserID, http.StatusSeeOther)
}

func (app *WebApp[T]) GetAuthLogin(w http.ResponseWriter, r *http.Request) {
	template.Must(template.New("login").Parse(loginTmpl)).Execute(w, app)
}

func (app *WebApp[T]) PostAuthSendLoginCode(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	err := app.AuthHandler().SendLoginCode(phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "UserID",
		Value: phone,
	})
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (app *WebApp[T]) GetAuthSendLoginCode(w http.ResponseWriter, r *http.Request) {
	template.Must(template.New("send-login-code").Parse(sendLoginCodeTmpl)).Execute(w, app)
}
