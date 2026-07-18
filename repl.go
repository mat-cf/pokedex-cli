package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mat-cf/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	NextURL *string
	PreviousURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			if !scanner.Scan() {
				break
			}

			input := cleanInput(scanner.Text())
			if len(input) == 0 {
				continue
			}
		
			cmd, ok := getCommands()[input[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			cmd.callback(cfg)
		}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Map the locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Map the previous locations",
			callback: commandMapB,
		},
	}
}