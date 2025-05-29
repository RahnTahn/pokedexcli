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
		fmt.Printf("Your command was: %v", answer[0])
		fmt.Println()

	}
}
func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(text))
	return cleaned
}
