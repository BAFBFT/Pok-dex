package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, param string) error {
	if param != "" {
		return errors.New("map command accepts no parameters")
	}
	pageURL := ""
	if cfg.nextLocationsURL != nil {
		pageURL = *cfg.nextLocationsURL
	} else if cfg.prevLocationsURL != nil {
		fmt.Println("you're on the last page")
		return nil
	}
	return showLocations(cfg, pageURL)
}

func commandMapB(cfg *config, param string) error {
	if param != "" {
		return errors.New("mapb command accepts no parameters")
	}
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return showLocations(cfg, *cfg.prevLocationsURL)
}

func showLocations(cfg *config, pageURL string) error {
	locations, err := cfg.pokeapiClient.ListLocations(pageURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
