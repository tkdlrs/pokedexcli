package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scannerThing := bufio.NewScanner(os.Stdin)
	//
	fmt.Print("Pokedex > ")
	for scannerThing.Scan() {
		usersInput := scannerThing.Text()
		//
		fmt.Println("Your command was:", cleanInput(usersInput)[0])
		//
		fmt.Print("Pokedex > ")
	}
	if err := scannerThing.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
