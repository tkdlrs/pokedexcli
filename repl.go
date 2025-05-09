package main

import "strings"

func cleanInput(text string) []string {
	var wordsSlice []string
	//
	removeWhiteSpace := strings.TrimSpace(text)
	removeWhiteSpace = strings.ReplaceAll(removeWhiteSpace, "  ", " ")
	lowercase := strings.ToLower(removeWhiteSpace)
	splitWords := strings.Split(lowercase, " ")
	wordsSlice = splitWords
	//
	return wordsSlice
}
