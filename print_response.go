package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func PrintResponse(resp *http.Response) {
	fmt.Println(resp.Status)
	lastByte := byte('\n')
	for {
		var buf [1]byte
		n, err := resp.Body.Read(buf[:])
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		if n != 1 {
			panic(n)
		}
		lastByte = buf[0]
		fmt.Printf("%s", string(lastByte))
	}
	if lastByte != '\n' {
		fmt.Println()
	}
}
