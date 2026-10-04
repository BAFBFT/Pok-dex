package main

import (
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
}

// create as a function to avoid, initialisation cycle
// if it was a variable, help command would need commandHelp function to be initialised
// commandHelp function would need to read var commands to be initialised
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
	}
}

func cleanInput(text string) []string {
	var result []string

	trimmed := strings.TrimSpace(text)
	i := strings.Index(trimmed, " ")
	for i > -1 {
		result = append(result, strings.ToLower(trimmed[:i]))
		trimmed = trimmed[i+1:]
		i = strings.Index(trimmed, " ")
	}
	return append(result, strings.ToLower(trimmed))
}
