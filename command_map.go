package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func getPokeLocations(url string) (pokeLocations, error) {
	res, err := http.Get(url)
	if err != nil {
		return pokeLocations{}, err
	}
	// body is type []byte, need to unmarshal it
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
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

func commandMap(cfg *config) error {
	locations, err := getPokeLocations("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}

	// update config
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
