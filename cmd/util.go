package cmd

import (
	"os"

	"golang.org/x/term"
)

func isValidFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func readPassword() []byte {
	password, _ := term.ReadPassword(0)
	return password
}
