package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex! \n Usage:")
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Exiting program")
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "help with the pokedex",
			callback:    commandHelp,
		},
	}
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		answer := cleanInput(reader.Text())
		value, exists := commands[answer[0]]
		if exists {
			value.callback()
		} else {
			fmt.Print("Unknown command")
		}
		fmt.Println()

	}
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(text))
	return cleaned
}
