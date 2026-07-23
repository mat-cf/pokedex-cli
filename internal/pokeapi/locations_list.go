package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("cache hit!")
		locations := LocationAreaResponse{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locations, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locations := LocationAreaResponse{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, data)

	return locations, nil
}