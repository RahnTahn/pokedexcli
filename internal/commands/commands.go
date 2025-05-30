package commands

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var Commands map[string]cliCommand

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex! \n Usage:")
	for _, value := range Commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Exiting program")
}

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "help with the pokedex",
			callback:    CommandHelp,
		},
	}
}
