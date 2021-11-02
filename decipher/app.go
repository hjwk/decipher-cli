package decipher

import (
	"flag"
	"fmt"
	"os"
)

type Command interface {
	FromArgs([]string) error
	Run() error
	Name() string
}

// CLI runs the decipher command line app and returns its exit status.
func CLI(args []string) int {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "You need to use a subcommand\n")
		return 1
	}

	commands := []Command{
		NewCipherCommand(),
		NewDecipherCommand(),
	}

	command := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name() == command {
			cmd.FromArgs(os.Args[2:])
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
				return 1
			}

			return 0
		}
	}

	return 0
}

type DecipherCommand struct {
	flagset    *flag.FlagSet
	cipher     string
	ciphertext string
	lang       string
}

func NewDecipherCommand() *DecipherCommand {
	cmd := DecipherCommand{flagset: flag.NewFlagSet("decipher", flag.ContinueOnError)}

	cmd.flagset.StringVar(&cmd.cipher, "c", "caesar", "Cipher algorithm")
	cmd.flagset.StringVar(&cmd.ciphertext, "t", "", "Ciphertext")
	cmd.flagset.StringVar(&cmd.lang, "l", "eng", "Supposed language of the ciphertext")

	return &cmd
}

func (cmd *DecipherCommand) FromArgs(args []string) error {
	if err := cmd.flagset.Parse(args); err != nil {
		return err
	}

	return nil
}

func (cmd *DecipherCommand) Run() error {
	answers := breakCaesar(cmd.ciphertext, cmd.lang)

	for _, solution := range answers {
		fmt.Println(solution)
	}

	return nil
}

func (cmd *DecipherCommand) Name() string {
	return cmd.flagset.Name()
}

type CipherCommand struct {
	flagset   *flag.FlagSet
	cipher    string
	plaintext string
	shift     int
}

func NewCipherCommand() *CipherCommand {
	cmd := CipherCommand{flagset: flag.NewFlagSet("cipher", flag.ContinueOnError)}

	cmd.flagset.StringVar(&cmd.cipher, "c", "caesar", "Cipher")
	cmd.flagset.StringVar(&cmd.plaintext, "t", "", "Plaintext")
	cmd.flagset.IntVar(&cmd.shift, "s", 3, "Shift")

	return &cmd
}

func (cmd *CipherCommand) FromArgs(args []string) error {
	if err := cmd.flagset.Parse(args); err != nil {
		return err
	}

	return nil
}

func (cmd *CipherCommand) Run() error {
	answer := caesar(cmd.plaintext, cmd.shift)

	_, err := fmt.Print(answer)
	return err
}

func (cmd *CipherCommand) Name() string {
	return cmd.flagset.Name()
}
