package cmd

import (
	"a-lafon/go-file-encrypt/filecrypt"
	"fmt"
	"log"

	"github.com/fatih/color"
)

func decrypt(args []string) {
	if len(args) < 3 {
		log.Fatalln("file path must be specified. For more info run help command")
	}

	file := args[2]

	if !isValidFile(file) {
		log.Fatalln("file not found")
	}
	fmt.Print("\nEnter password: ")
	password := readPassword()

	fmt.Println("\n\nDecrypting ...")

	err := filecrypt.Decrypt(file, password)
	if err != nil {
		log.Fatalln("error on file decryption", err)
	}

	color.Green("\nFile sucessfully decrypted")
}
