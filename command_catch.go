package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	// Get the information on the Pokemon.
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	//
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	// Determine a bool as to weather the pokemon was captured or not
	caughtPokemon := catchPokemon(pokemon.BaseExperience)
	if !caughtPokemon {
		fmt.Printf("%s escaped!\n", name)
		return nil
	} else {
		fmt.Printf("%s was caught!\n", name)
		// keep a list of caught pokemon
		cfg.caughtPokemon[name] = pokemon
		return nil
	}
}

func catchPokemon(baseExperience int) bool {
	randomNumber := rand.Intn(99) + 1
	return randomNumber < baseExperience
}
