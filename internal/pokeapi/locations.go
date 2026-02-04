package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(listURL *string) (apiRes, error) {
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

	data := apiRes{}
	err = json.Unmarshal(body, &data)
	return data, nil
}

type apiRes struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
