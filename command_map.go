package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	// http get https://pokeapi.co/api/v2/location-area/
	data := pokeapiGet("https://pokeapi.co/api/v2/location-area/")
	locations := data.Results
	for _, loc := range locations {
		fmt.Printf("name: %s\n", loc.Name)
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
