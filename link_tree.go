package util

import (
	"encoding/json"
	"net/http"
)

type LinkTree struct {
	Head  *HTMLHead
	Links []Link
}

func (tree *LinkTree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tree)
}
