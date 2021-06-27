package decipher

import (
	"flag"
	"fmt"
	"os"
)

// CLI runs the decipher command line app and returns its exit status.
func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	cipher    string
	plaintext string
	shift     int
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("cipher", flag.ContinueOnError)
	fl.StringVar(&app.cipher, "c", "caesar", "Cipher")
	fl.StringVar(&app.plaintext, "t", "", "Plaintext")
	fl.IntVar(&app.shift, "s", 3, "Shift")

	if err := fl.Parse(args); err != nil {
		return err
	}

	return nil
}

func (app *appEnv) run() error {
	answer := caesar(app.plaintext, app.shift)

	_, err := fmt.Print(answer)
	return err
}
