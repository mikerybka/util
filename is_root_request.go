package util

import "net/http"

func IsRootRequest(r *http.Request) bool {
	return len(ParsePath(r.URL.Path)) == 0
}
