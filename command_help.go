package main

import (
	"fmt"
	"maps"
	"slices"
)

func commandHelp(cfg *config) error {
	names := slices.Sorted(maps.Keys(cfg.commands))
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, name := range names {
		fmt.Printf("%s: %s\n", name, cfg.commands[name].description)
	}
	return nil
}
