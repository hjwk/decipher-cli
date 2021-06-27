package main

import (
	"os"

	"github.com/hjwk/decipher/decipher"
)

func main() {
	os.Exit(decipher.CLI(os.Args[1:]))
}
