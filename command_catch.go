package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, param string) error {
	if param == "" {
		return errors.New("enter a valid pokemon name")
	}
	pokemonInfo, err := cfg.pokeapiClient.GetPokemon(param)
	if err != nil {
		return err
	}

	// number of tries needed to catch pokemon
	fmt.Println("Throwing a Pokeball at " + param + "...")

	_, ok := cfg.pokemonCatchCounter[param]
	if !ok {
		baseExperience := pokemonInfo.BaseExperience
		tries := baseExperience / 50
		cfg.pokemonCatchCounter[param] = tries
	}

	if cfg.pokemonCatchCounter[param] == 0 {
		fmt.Println(param + " was caught!")
		cfg.pokedex[param] = pokemonInfo
	} else {
		cfg.pokemonCatchCounter[param] = cfg.pokemonCatchCounter[param] - 1
		fmt.Println(param + " escaped!")
	}

	return nil
}
