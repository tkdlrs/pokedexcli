package main

import "fmt"

func commandHelp() error {
	outputString := "help: Displays a help message\nexit: Exit the Pokedex\n"
	fmt.Printf("Welcome to the Pokedex! \nUsage: \n\n%v", outputString)
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
