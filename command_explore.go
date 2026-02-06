package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("No location provided to explore")
	}

	fmt.Printf("%s\n", params[0])
	return nil
	/*
		data, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
		if err != nil {
			return errors.New("temp")
		}
		locations := data.Results

		for _, loc := range locations {
			fmt.Printf("name: %s\n", loc.Name)
		}

		cfg.nextLocationsURL = data.Next
		cfg.prevLocationsURL = data.Previous

		return nil
	*/
}
