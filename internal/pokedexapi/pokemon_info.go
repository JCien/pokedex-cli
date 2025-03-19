package pokedexapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(PokemonName string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + PokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonInfo := PokemonInfo{}
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, data)
	return pokemonInfo, nil
}
