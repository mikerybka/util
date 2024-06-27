package util

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type App struct {
	Dir string
}

func (a *App) Handler(r *http.Request) http.Handler {
	switch a.Kind() {
	case "PingServer":
		return &PingServer{}
	case "SchemaCafe":
		return &SchemaCafe{}
	case "LinkTree":
		return &LinkTree{}
	default:
		return http.NotFoundHandler()
	}
}

func (a *App) Config() *AppConfig {
	path := filepath.Join(a.Dir, "config.json")
	return ReadJSONFile[*AppConfig](path)
}

func (a *App) Kind() string {
	return a.Config().Kind
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := a.Handler(r)

	path := filepath.Join(a.Dir, "data.json")
	b, _ := os.ReadFile(path)
	json.Unmarshal(b, h)

	h.ServeHTTP(w, r)

	if IsMutation(r) {
		b, err := json.MarshalIndent(h, "", "  ")
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(path, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}