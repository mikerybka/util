package util

import (
	"fmt"
	"net/http"
)

func ServeMyIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, ClientIP(r))
}
