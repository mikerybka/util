package util

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetPasswordFromStdin(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // move to next line after input
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}
	return string(bytePassword), nil
}
