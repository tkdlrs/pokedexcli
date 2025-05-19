package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tkdlrs/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	capturedPokemon  map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	cfg.capturedPokemon = make(map[string]pokeapi.Pokemon)
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		//
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		//
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		//
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("Captured pokemon include: ", cfg.capturedPokemon)
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemans.",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Takes a parameter of a location area and returns a list of the Pokeman's in that area",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas from the PokeAPI",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas from the PokeAPI",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
