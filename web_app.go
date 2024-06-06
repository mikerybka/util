package util

import (
	"fmt"
	"net/http"
	"strings"
)

type WebApp struct {
	Name         string
	Description  string
	Author       string
	Keywords     []string
	Favicon      []byte
	Types        map[string]Type
	RootType     string
	TwilioClient *TwilioClient
	Files        FileSystem
}

func (app *WebApp) Validate() error {
	if _, ok := app.Types["App"]; ok {
		return fmt.Errorf("type with name `App` not allowed")
	}
	return nil
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api") {
		types := app.Types
		types["App"] = Type{
			IsMap:    true,
			ElemType: app.RootType,
		}
		http.StripPrefix("/api", &MultiUserApp{
			Twilio:    app.TwilioClient,
			AuthFiles: app.Files.Dig("auth"),
			App: &WebAPI{
				Types:    app.Types,
				RootType: "App",
				Data:     app.Files.Dig("data"),
			},
		}).ServeHTTP(w, r)
		return
	}

	app.Frontend().ServeHTTP(w, r)
}

func (app *WebApp) Frontend() *WebFrontend {
	return &WebFrontend{
		Favicon:      app.Favicon,
		RootTitle:    app.Name,
		MetaDesc:     app.Description,
		MetaAuthor:   app.Author,
		MetaKeywords: app.Keywords,
	}
}
