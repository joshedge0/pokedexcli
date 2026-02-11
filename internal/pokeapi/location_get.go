package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName *string) (ResLocationDeep, error) {
	url := baseURL + "/location-area/" + *locationName

	val, exists := c.cache.Get(url)
	if exists {
		data := ResLocationDeep{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return ResLocationDeep{}, err
		}
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return ResLocationDeep{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return ResLocationDeep{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return ResLocationDeep{}, err
	}

	data := ResLocationDeep{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ResLocationDeep{}, err
	}

	c.cache.Add(url, body)
	return data, nil
}
