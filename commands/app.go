package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/hjwk/decipher/caesar"
)

type Command interface {
	FromArgs([]string) error
	Run() error
	Name() string
}

// CLI runs the decipher command line app and returns its exit status.
func CLI(args []string) int {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "You need to use a subcommand. Supported subcommands are: \n- cipher\n- decipher\n")
		return 1
	}

	commands := []Command{
		NewCipherCommand(),
		NewDecipherCommand(),
	}

	command := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name() == command {
			err := cmd.FromArgs(os.Args[2:])
			if err != nil {
				return 1
			}

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

	if !isCipherSupported(cmd.cipher) {
		cmd.flagset.Usage()
		return flag.ErrHelp
	}

	if !isLangSupported(cmd.lang) {
		cmd.flagset.Usage()
		return flag.ErrHelp
	}

	return nil
}

func (cmd *DecipherCommand) Run() error {
	shift, deciphered := caesar.Decipher(cmd.ciphertext, cmd.lang)

	fmt.Printf("Shift: %d\nDeciphered message: %s\n", shift, deciphered)

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

	if !isCipherSupported(cmd.cipher) {
		cmd.flagset.Usage()
		return flag.ErrHelp
	}

	return nil
}

func (cmd *CipherCommand) Run() error {
	answer := caesar.Encipher(cmd.plaintext, cmd.shift)

	_, err := fmt.Print(answer)
	return err
}

func (cmd *CipherCommand) Name() string {
	return cmd.flagset.Name()
}

func isCipherSupported(cipher string) bool {
	switch cipher {
	case "casear", "scytale":
		return true
	default:
		return false
	}
}

func isLangSupported(lang string) bool {
	switch lang {
	case "fr", "eng":
		return true
	default:
		return false
	}
}
