package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("No location provided to explore")
	}

	fmt.Printf("Exploring %s...\n", params[0])

	data, err := cfg.pokeapiClient.GetLocation(&params[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	pokemon := data.PokemonEncounters

	for _, poke := range pokemon {
		fmt.Printf("- %s\n", poke.Pokemon.Name)
	}

	return nil
}
