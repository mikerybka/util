package util

import "net/http"

type SingleUserAuthApp struct {
	Workdir string
	App     http.Handler
}
