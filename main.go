package main

import (
	"time"

	"github.com/joshedge0/pokedexcli/internal/pokeapi"
	"github.com/joshedge0/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokecache:     cache,
	}

	startRepl(cfg)
}
