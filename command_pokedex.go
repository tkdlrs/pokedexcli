package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Print("\nYour Pokedex:\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	fmt.Println()
	return nil
}
