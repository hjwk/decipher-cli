package main

import (
	"os"

	"github.com/hjwk/decipher-cli/commands"
)

func main() {
	os.Exit(commands.CLI(os.Args[1:]))
}
