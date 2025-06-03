package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type locations struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}
type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var commands map[string]cliCommand
var mapCurrent string

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex! \n Usage:")
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}
	return nil
}
func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Exiting program")
}

func CommandMap() error {

	jsonData, _ := jsonGrabber(mapCurrent)
	var locationList locations
	err := json.Unmarshal(jsonData, &locationList)
	if err != nil {
		return errors.New("json failed")
	}

	mapCurrent = locationList.Next

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandMapb() error {
	return nil
}

func init() {
	mapCurrent = "https://pokeapi.co/api/v2/location?limit=20"
	commands = map[string]cliCommand{
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
		"map": {
			Name:        "map",
			Description: "locations in pokemon",
			Callback:    CommandMap,
		},
	}

}

func GetCommands() map[string]cliCommand {
	return commands
}

func jsonGrabber(url string) ([]byte, error) {
	res, err := http.Get(mapCurrent)
	if err != nil {
		return nil, errors.New("res failed")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("body failed")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.New("connection failed")
	}

	return body, nil
}
