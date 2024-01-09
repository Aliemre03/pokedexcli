package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (*Pokemon, error) {
	var pokemon Pokemon
	url := baseURL + "pokemon/" + pokemonName
	data, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return &Pokemon{}, err
		}

		return &pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return &Pokemon{}, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return &Pokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return &Pokemon{}, err
	}

	c.cache.Add(url, data)

	return &pokemon, nil
}
