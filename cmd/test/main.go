package main

import (
	"fmt"
	"net/http"
)

func IntToID(v int64) string {
	// Get the hex encoding
	hex := fmt.Sprintf("%016x", v)

	// Swap out our custom values and reverse the output
	mapping := map[rune]string{
		'0': "k",
		'1': "e",
		'2': "g",
		'3': "n",
		'4': "d",
		'5': "r",
		'6': "f",
		'7': "h",
		'8': "b",
		'9': "s",
		'a': "j",
		'b': "m",
		'c': "p",
		'd': "q",
		'e': "c",
		'f': "a",
	}
	res := ""
	for _, r := range hex {
		res = mapping[r] + res
	}
	return res
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := r.Header.Get("Accept")
		w.Write([]byte(s))
	})
	http.ListenAndServe(":8000", nil)
}
