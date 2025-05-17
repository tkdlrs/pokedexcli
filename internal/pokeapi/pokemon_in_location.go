package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListPokemonInLocation -
func (c *Client) ListPokemonInLocation(areaToExplore string) (RespDeepLocation, error) {
	if areaToExplore == "" {
		return RespDeepLocation{}, fmt.Errorf("ERROR: area to explore not provided")
	}
	url := baseURL + "/location-area/" + areaToExplore
	//
	if val, ok := c.cache.Get(url); ok {
		completeLocationResp := RespDeepLocation{}
		err := json.Unmarshal(val, &completeLocationResp)
		if err != nil {
			return RespDeepLocation{}, err
		}
		return completeLocationResp, nil
	}
	//
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}
	//
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close()
	//
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}
	//
	completeLocationResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &completeLocationResp)
	if err != nil {
		return RespDeepLocation{}, err
	}
	//
	return completeLocationResp, nil
}
