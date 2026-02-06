package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(listURL *string) (resLocationShallow, error) {
	url := baseURL + "/location-area"
	if listURL != nil {
		url = *listURL
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	data := resLocationShallow{}
	err = json.Unmarshal(body, &data)
	return data, nil
}

func (c *Client) GetLocation(locationName *string) (resLocationDeep, error) {
	url := baseURL + "/location-area/" + *locationName

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	data := resLocationDeep{}
	err = json.Unmarshal(body, &data)
	return data, nil
}

type resLocationShallow struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type resLocationDeep struct {
	//EncounterMethodRates string   `json:"encounter_method_rates"`
	//GameIndex            int      `json:"game_index"`
	//Id                   int      `json:"id"`
	Location location `json:"location"`
	Name     string   `json:"name"`
	//Names                []string `json:"names"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
