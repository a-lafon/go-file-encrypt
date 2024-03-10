package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	command := os.Args[1]
	fmt.Println(command)
	switch command {
	case "encrypt":
		execEncrypt()
	case "decrypt":
		execDecrypt()
	default:
		printHelp()
	}
}

func execEncrypt() {
	if len(os.Args) < 3 {
		log.Fatalln("file path must be specified. For more info run help command")
	}

	file := os.Args[2]

	if !isValidFile(file) {
		log.Fatalln("file not found")
	}
	fmt.Println(file)

	password := getPassword()
	fmt.Println(password)
}

func isValidFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func getPassword() []byte {
	fmt.Print("\nEnter password: ")
	password, _ := term.ReadPassword(0)

	isValid, err := isValidPassword(string(password))
	if !isValid {
		color.Yellow("\n%s\n", err)
		return getPassword()
	}

	fmt.Print("\nConfirm password: ")
	passwordConfirm, _ := term.ReadPassword(0)

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

func execDecrypt() {

}

func printHelp() {
	fmt.Println("Simpe file encrypter")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\t go run . [command]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypt file given a password")
	fmt.Println("\t decrypt\tDecrypt file using a password")
	fmt.Println("\t help\t\tDisplays help text")
}
