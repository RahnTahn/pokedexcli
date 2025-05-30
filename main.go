package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		answer := cleanInput(reader.Text())
		value, exists := Commands[answer[0]]
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
