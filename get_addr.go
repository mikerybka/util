package util

import "os"

func GetAddr() string {
	addr := ":3000"
	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	}
	return addr
}
