package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(listURL *string) (ResLocationShallow, error) {
	url := baseURL + "/location-area"
	if listURL != nil {
		url = *listURL
	}

	val, exists := c.cache.Get(url)
	if exists {
		data := ResLocationShallow{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return ResLocationShallow{}, err
		}
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return ResLocationShallow{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return ResLocationShallow{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return ResLocationShallow{}, err
	}
	c.cache.Add(url, body)

	data := ResLocationShallow{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ResLocationShallow{}, err
	}

	return data, nil
}
