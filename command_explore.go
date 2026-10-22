package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, param string) error {
	if param == "" {
		return errors.New("must provide an area as name or id")
	}
	fmt.Println("Exploring " + param + "...")
	areaInfo, err := cfg.pokeapiClient.GetLocationArea(param)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, loc := range areaInfo.PokemonEncounters {
		fmt.Println("- " + loc.Pokemon.Name)
	}
	return nil
}
