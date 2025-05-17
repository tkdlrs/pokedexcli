package main

import (
	"fmt"
)

func commandExplore(cfg *config, explorableArea string) error {
	fmt.Println("Exploring " + explorableArea + "...")
	//
	deepLocationResp, err := cfg.pokeapiClient.ListPokemonInLocation(explorableArea)
	if err != nil {
		return err
	}
	// fmt.Println(deepLocationResp)
	// filter out the un-necessary information
	pokemons := []string{}
	for _, encounters := range deepLocationResp.PokemonEncounters {
		pokemons = append(pokemons, encounters.Pokemon.Name)
	}
	//
	fmt.Println("Found Pokemon:")
	for _, name := range pokemons {
		fmt.Printf(" - %v\n", name)
	}
	//
	return nil
}
