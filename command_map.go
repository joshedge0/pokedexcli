package main

import (
	"fmt"
)

func commandMap() error {
	// http get https://pokeapi.co/api/v2/location-area/
	locations := pokeapiGet("https://pokeapi.co/api/v2/location-area/")
	for _, loc := range locations {
		fmt.Printf("name: %s\n", loc.Name)
	}
	return nil
}
