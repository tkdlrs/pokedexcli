package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
		// fmt.Printf("Your command was: %s\n", commandName)
		if thing, ok := supportedCommands[commandName]; ok {
			thing.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
