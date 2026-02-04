package main

import "log"

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	startRepl(cfg)
}

func loadConfig() (*config, error) {
	return &config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}, nil
}

type config struct {
	Next     string
	Previous string
}
