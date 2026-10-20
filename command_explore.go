package main

import (
	"errors"
	"fmt"
)

// Should probably have some error handling logic here, to handle not found requests
func commandExplore(cfg *config, param string) error {
	if param == "" {
		return errors.New("must provide an area as name or id")
	}
	areaInfo, err := explorePokeLocations(cfg.locationAreaEndpoint + param)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + param + "...")
	fmt.Println("Found Pokemon:")
	// iterate ove the names
	for _, loc := range areaInfo.PokemonEncounters {
		fmt.Println("- " + loc.Pokemon.Name)
	}
	return nil
}
