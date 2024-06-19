package util

import (
	"net/http"
)

func ServeMethod(m string, v any, w http.ResponseWriter, r *http.Request) {
	// t := reflect.ValueOf(v)
	// m = t.MethodByName(m)
	// if r.Method == "GET" {

	// }
}
