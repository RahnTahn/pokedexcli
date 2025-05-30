package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/commands"
	"strings"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		answer := cleanInput(reader.Text())
		cmds := commands.GetCommands()
		value, exists := cmds[answer[0]]

		if len(answer) == 0 {
			continue
		}

		if exists {
			value.Callback()
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
