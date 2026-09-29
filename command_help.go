package main

import (
	"fmt"
	"maps"
	"slices"
)

func commandHelp() error {
	cmds := getCommands()
	names := slices.Sorted(maps.Keys(cmds))
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, name := range names {
		fmt.Printf("%s: %s\n", name, cmds[name].description)
	}
	return nil
}
