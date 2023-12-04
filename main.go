package util

import (
	"net/http"
	"os"
)

func Main(s http.Handler) {
	path := os.Args[1]
	port := os.Args[2]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.ServeHTTP(w, r)
		if IsMutation(r) {
			err := WriteJSONFile(path, s)
			if err != nil {
				panic(err)
			}
		}
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
