package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, param string) error
}

type config struct {
	pokeapiClient       *pokeapi.Client
	pokedex             map[string]pokeapi.Pokemon
	pokemonCatchCounter map[string]int
	nextLocationsURL    *string
	prevLocationsURL    *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			// EOF (Ctrl-D) or read error
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			return
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		param := ""
		if len(words) > 1 {
			param = words[1]
		}

		if err := cmd.callback(cfg, param); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// a function rather than a variable to avoid an initialisation cycle:
// commandHelp reads the commands, and the commands reference commandHelp
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
			description: "Displays the names of the next 20 location areas in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon World",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of all Pokémon found in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View details of caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all pokemon in the pokedex",
			callback:    commandPokedex,
		},
	}
}
