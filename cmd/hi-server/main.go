package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi")
	})
	port := os.Getenv("PORT")
	addr := ":" + port
	http.ListenAndServe(addr, nil)
}
