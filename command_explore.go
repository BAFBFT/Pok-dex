package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, param string) error {
	if param == "" {
		errors.New("must provide an area as name or id")
	}
	areaInfo, err := explorePokeLocations(cfg.locationAreaEndpoint + param)
	if err != nil {
		return err
	}
	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	// iterate ove the names
	for _, loc := range areaInfo.PokemonEncounters {
		fmt.Println("- " + loc.Pokemon.Name)
	}
	return nil
}
