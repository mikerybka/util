package util

import (
	_ "embed"
	"errors"
	"net/http"
	"os"
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

func (app *WebApp[T]) HomeTmpl() *template.Template {
	return template.Must(template.New("home").Parse(homeTmpl))
}

func (app *WebApp[T]) OrgTmpl() *template.Template {
	return template.Must(template.New("org").Parse(orgTmpl))
}

func (app *WebApp[T]) GetRoot(w http.ResponseWriter, r *http.Request) {
	app.HomeTmpl().Execute(w, app)
}

func (app *WebApp[T]) PostRoot(w http.ResponseWriter, r *http.Request) {
	// TODO: handle admin features
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
	app.OrgTmpl().Execute(w, entries)
}

func (app *WebApp[T]) PutOrg(w http.ResponseWriter, r *http.Request) {

}
func (app *WebApp[T]) PostOrg(w http.ResponseWriter, r *http.Request) {

}
func (app *WebApp[T]) GetPath(w http.ResponseWriter, r *http.Request) {

}
func (app *WebApp[T]) PutPath(w http.ResponseWriter, r *http.Request)    {}
func (app *WebApp[T]) PostPath(w http.ResponseWriter, r *http.Request)   {}
func (app *WebApp[T]) PatchPath(w http.ResponseWriter, r *http.Request)  {}
func (app *WebApp[T]) DeletePath(w http.ResponseWriter, r *http.Request) {}
func (app *WebApp[T]) GetMeta(w http.ResponseWriter, r *http.Request)    {}
func (app *WebApp[T]) GetFavicon(w http.ResponseWriter, r *http.Request) {
	w.Write(app.Favicon)
}
func (app *WebApp[T]) PostAuthLogin(w http.ResponseWriter, r *http.Request)         {}
func (app *WebApp[T]) GetAuthLogin(w http.ResponseWriter, r *http.Request)          {}
func (app *WebApp[T]) PostAuthSendLoginCode(w http.ResponseWriter, r *http.Request) {}
func (app *WebApp[T]) GetAuthSendLoginCode(w http.ResponseWriter, r *http.Request)  {}
