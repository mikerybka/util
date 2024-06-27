package util

import "net/http"

type LinkTree struct {
	Head  *HTMLHead
	Links []Link
}

func (tree *LinkTree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	doc := &HTMLDocument{
		Head: tree.Head,
	}
	doc.ServeHTTP(w, r)
}
