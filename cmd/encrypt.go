package cmd

import (
	"a-lafon/go-file-encrypt/filecrypt"
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/fatih/color"
)

func encrypt(args []string) {
	if len(args) < 3 {
		log.Fatalln("file path must be specified. For more info run help command")
	}

	file := args[2]

	if !isValidFile(file) {
		log.Fatalln("file not found")
	}

	password := getPassword()
	fmt.Println("\n\nEncrypting ...")

	err := filecrypt.Encrypt(file, password)
	if err != nil {
		log.Fatalln("error on file encryption", err)
	}

	color.Green("\nFile sucessfully encrypted")
}

func getPassword() []byte {
	fmt.Print("\nEnter password: ")
	password := readPassword()

	isValid, err := isValidPassword(string(password))
	if !isValid {
		color.Yellow("\n%s\n", err)
		return getPassword()
	}

	fmt.Print("\nConfirm password: ")
	passwordConfirm := readPassword()

	if !isPasswordMatching(password, passwordConfirm) {
		color.Red("\npasswords don't match. Please try again\n")
		return getPassword()
	}

	return password
}

func isValidPassword(password string) (bool, error) {
	if len(password) < 5 {
		return false, errors.New("password must be at least 5 characters")
	}
	return true, nil
}

func isPasswordMatching(password []byte, password2 []byte) bool {
	return bytes.Equal(password, password2)
}
