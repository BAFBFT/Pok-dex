package main

import (
	"time"

	"pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient:       pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokedex:             make(map[string]pokeapi.Pokemon),
		pokemonCatchCounter: make(map[string]int),
	}
	startRepl(cfg)
}
