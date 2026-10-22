package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client fetches data from the PokeAPI, caching every response body by URL.
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

// ListLocations returns a page of location areas.
// An empty pageURL fetches the first page.
func (c *Client) ListLocations(pageURL string) (LocationAreas, error) {
	url := pageURL
	if url == "" {
		url = baseURL + "/location-area/"
	}
	var locations LocationAreas
	err := c.getJSON(url, &locations)
	return locations, err
}

// GetLocationArea returns a location area by name or id.
func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	var area LocationArea
	err := c.getJSON(baseURL+"/location-area/"+name, &area)
	return area, err
}

// GetPokemon returns a pokemon by name or id.
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var pokemon Pokemon
	err := c.getJSON(baseURL+"/pokemon/"+name, &pokemon)
	return pokemon, err
}

func (c *Client) getJSON(url string, v any) error {
	body, err := c.getBody(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (c *Client) getBody(url string) ([]byte, error) {
	if body, ok := c.cache.Get(url); ok {
		return body, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("not found")
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("request failed with status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// only cache successful responses
	c.cache.Add(url, body)
	return body, nil
}
