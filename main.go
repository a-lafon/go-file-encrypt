package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	command := os.Args[1]
	fmt.Println(command)
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
