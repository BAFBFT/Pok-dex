package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, param string) error {
	if param != "" {
		return errors.New("map command accepts no parameters")
	}
	locations, err := getPokeLocations(cfg.next)
	if err != nil {
		return err
	}
	// update config
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config, param string) error {
	if param != "" {
		return errors.New("mapb command accepts no parameters")
	}
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := getPokeLocations(cfg.previous)
	if err != nil {
		return err
	}
	// update config
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
