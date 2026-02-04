package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	data, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	locations := data.Results

	for _, loc := range locations {
		fmt.Printf("name: %s\n", loc.Name)
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New(`You're on the first page`)
	}

	data, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	locations := data.Results

	for _, loc := range locations {
		fmt.Printf("name: %s\n", loc.Name)
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	return nil
}
