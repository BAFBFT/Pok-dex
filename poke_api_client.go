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

func getPokeLocations(url string) (pokeLocations, error) {
	value, ok := cache.Get(url)
	var body []byte
	if ok {
		body = value
	} else {
		res, err := http.Get(url)
		if err != nil {
			return pokeLocations{}, err
		}
		// body is type []byte, need to unmarshal it
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		if err != nil {
			return pokeLocations{}, err
		}

		//store in cache if no errors
		cache.Add(url, body)
	}

	// unmarshal into pokeLocations struct
	var locations pokeLocations
	if err := json.Unmarshal(body, &locations); err != nil {
		return pokeLocations{}, err
	}
	return locations, nil
}
