package main

import (
	"errors"
	"fmt"
	"maps"
	"slices"
)

func commandPokedex(cfg *config, param string) error {
	if param != "" {
		return errors.New("pokedex command accepts no parameters")
	}

	fmt.Println("Your Pokedex:")
	for _, name := range slices.Sorted(maps.Keys(cfg.pokedex)) {
		fmt.Printf(" -%s\n", name)
	}
	return nil
}
