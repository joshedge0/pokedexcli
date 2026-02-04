package main

import "fmt"

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	data := pokeapiGet(cfg.Previous)
	locations := data.Results
	for _, loc := range locations {
		fmt.Printf("name: %s\n", loc.Name)
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	return nil
}
