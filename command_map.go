package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
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
