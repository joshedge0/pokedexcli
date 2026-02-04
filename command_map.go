package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.Next == "" {
		fmt.Println("You're on the last page")
		return nil
	}

	data := pokeapiGet(cfg.Next)
	locations := data.Results

	for _, loc := range locations {
		fmt.Printf("name: %s\n", loc.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	return nil
}
