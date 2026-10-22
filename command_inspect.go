package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, param string) error {
	if param == "" {
		return errors.New("enter a valid pokemon name")
	}

	_, ok := cfg.pokedex[param]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	pokeinfo := cfg.pokedex[param]
	fmt.Printf("Name: %s\n", pokeinfo.Name)
	fmt.Printf("Height: %d\n", pokeinfo.Height)
	fmt.Printf("Weight: %d\n", pokeinfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokeinfo.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemontype := range pokeinfo.Pokemontypes {
		fmt.Printf("  - %s\n", pokemontype.Pokemontype.Typename)
	}
	return nil
}
