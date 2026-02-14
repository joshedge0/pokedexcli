package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joshedge0/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputString := cleanInput(scanner.Text())
		commandName := inputString[0]
		params := inputString[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, params)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	strs := strings.Fields(strings.ToLower(text))
	return strs
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 map locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a map location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Capture a Pokemon",
			callback:    commandCatch,
		},
	}
}
