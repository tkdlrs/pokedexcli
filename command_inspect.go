package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	// Check to see if this pokemon has been caught. Get the information on the Pokemon.
	name := args[0]
	pokeInfo, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	// Print info
	fmt.Printf("Name: %s \nHeight: %v \nWeight: %v \n", pokeInfo.Name, pokeInfo.Height, pokeInfo.Weight)
	// Print stats
	fmt.Println("Stats:")
	for _, pokeStat := range pokeInfo.Stats {
		fmt.Printf("  -%s: %v \n", pokeStat.Stat.Name, pokeStat.BaseStat)
	}
	// Print types
	fmt.Println("Types:")
	for _, pokeType := range pokeInfo.Types {
		fmt.Printf("  - %v \n", pokeType.Type.Name)
	}
	return nil
}
