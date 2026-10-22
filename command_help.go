package main

import (
	"errors"
	"fmt"
	"maps"
	"slices"
)

func commandHelp(cfg *config, param string) error {
	if param != "" {
		return errors.New("help command accepts no parameters")
	}
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, name := range slices.Sorted(maps.Keys(commands)) {
		fmt.Printf("%s: %s\n", name, commands[name].description)
	}
	return nil
}
