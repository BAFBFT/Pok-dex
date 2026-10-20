package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

var cache = pokecache.NewCache(300 * time.Second)

type pokeLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokeLocations struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []pokeLocation `json:"results"`
}

// only the names are needed
type locationArea struct {
	PokemonEncounters []struct { // slice of pokemons
		Pokemon struct { // struct including name field
			Name string `json:"name"` // name field
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func getUrlBody(url string) ([]byte, error) {
	value, ok := cache.Get(url)
	var body []byte
	if ok {
		body = value
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		// body is type []byte, need to unmarshal it
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		if err != nil {
			return nil, err
		}

		//store in cache if no errors
		cache.Add(url, body)
	}
	return body, nil
}

func getPokeLocations(url string) (pokeLocations, error) {
	body, err := getUrlBody(url)
	if err != nil {
		return pokeLocations{}, err
	}

	// unmarshal into pokeLocations struct
	var locations pokeLocations
	if err := json.Unmarshal(body, &locations); err != nil {
		return pokeLocations{}, err
	}
	return locations, nil
}

func explorePokeLocations(url string) (locationArea, error) {
	body, err := getUrlBody(url)
	if err != nil {
		return locationArea{}, err
	}

	// unmarshal into pokeLocations struct
	var areaInfo locationArea
	if err := json.Unmarshal(body, &areaInfo); err != nil {
		return locationArea{}, err
	}
	return areaInfo, nil
}

func getPokemonInfo(url string) (pokemon, error) {
	body, err := getUrlBody(url)
	if err != nil {
		return pokemon{}, err
	}

	// unmarshal into pokeLocations struct
	var pokemonInfo pokemon
	if err := json.Unmarshal(body, &pokemonInfo); err != nil {
		return pokemon{}, err
	}
	return pokemonInfo, nil
}
