package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	//
	configurations := config{
		Next:     "",
		Previous: "",
	}
	//
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
		//
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&configurations)
			if err != nil {
				fmt.Println(err)
			}
			//
			// fmt.Println(configurations.Next, "configurations.Next")
			// fmt.Println(configurations.Previous, "configurations.Previous")
			//
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
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas from the PokeAPI",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas from the PokeAPI",
			callback:    commandMapB,
		},
	}
}

type config struct {
	Next     string
	Previous string
}
