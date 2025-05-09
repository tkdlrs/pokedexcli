package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("error")
}

func commandHelp() error {
	outputString := "help: Displays a help message\nexit: Exit the Pokedex\n"
	fmt.Printf("Welcome to the Pokedex! \nUsage: \n\n%v", outputString)
	// for _, val := range supportedCommands {
	// 	outputString += fmt.Sprintf("%v: %v", val.name, val.description)
	// }
	return nil
}

var supportedCommands = map[string]cliCommand{
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
}
