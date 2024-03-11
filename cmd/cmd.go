package cmd

import (
	"os"

	"github.com/fatih/color"
)

func Exec(args []string) {
	if len(args) < 2 {
		printHelp()
		os.Exit(0)
	}

	command := args[1]

	switch command {
	case "encrypt":
		encrypt(args)
	case "decrypt":
		decrypt(args)
	case "help":
		printHelp()
	default:
		color.Yellow("Command not found\n\n")
		printHelp()
	}
}
