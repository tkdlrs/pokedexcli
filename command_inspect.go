package main

import (
	"errors"
	"fmt"

	"github.com/tkdlrs/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	// Check to see if this pokemon has been caught. Get the information on the Pokemon.
	name := args[0]
	//
	var pokeInfo pokeapi.Pokemon
	var ok bool
	pokeInfo, ok = cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		// Print info
		fmt.Printf("Name: %s \nHeight: %v \nWeight: %v \n", pokeInfo.Name, pokeInfo.Height, pokeInfo.Weight)
		// Print stats
		fmt.Println("Stats:")
		for _, pokeStat := range pokeInfo.Stats {
			fmt.Printf("  -%v: %v \n", pokeStat.Stat.Name, pokeStat.BaseStat)
		}
		// Print types
		fmt.Println("Types:")
		for _, pokeType := range pokeInfo.Types {
			fmt.Printf("  - %v \n", pokeType.Type.Name)
		}
		return nil
	}

}
