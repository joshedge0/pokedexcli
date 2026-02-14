package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("No Pokemon provided to capture")
	}

	// handle for if pokemon name not in options

	fmt.Printf("Throwing a Pokeball at %s\n", params[0])

	data, err := cfg.pokeapiClient.GetPokemon(&params[0])
	if err != nil {
		return err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randFloat := r.Float32()

	oddsOfCapture := randFloat - (0.001 * data.BaseExperience)

	if oddsOfCapture > 0.45 {
		fmt.Printf("%s was caught!\n", data.Name)
		return nil
	}

	fmt.Printf("%s got away!\n", data.Name)
	return nil
}
