package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName *string) (PokemonApiBody, error) {
	url := baseURL + "/pokemon/" + *pokemonName

	val, exists := c.cache.Get(url)
	if exists {
		data := PokemonApiBody{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return PokemonApiBody{}, err
		}
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonApiBody{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonApiBody{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokemonApiBody{}, err
	}

	data := PokemonApiBody{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return PokemonApiBody{}, err
	}

	c.cache.Add(url, body)
	return data, nil
}
