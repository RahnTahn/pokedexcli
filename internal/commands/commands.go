package commands

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var Commands map[string]cliCommand

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex! \n Usage:")
	for _, value := range Commands {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
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
			Name:        "exit",
			Description: "exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "help with the pokedex",
			Callback:    CommandHelp,
		},
	}
}

func GetCommands() map[string]cliCommand {
	return Commands
}
