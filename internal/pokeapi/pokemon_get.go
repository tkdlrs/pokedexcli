package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	/*	Caching...
		It does make sense to catch a pokemon's information.
		That way when a player fails to catch and tries again the second time will be faster.
	*/
	if pokeInfo, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(pokeInfo, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	//
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	//
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	//
	pokeInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	//
	pokemonResp := Pokemon{}
	err = json.Unmarshal(pokeInfo, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	//
	c.cache.Add(url, pokeInfo)
	//
	return pokemonResp, nil

}
