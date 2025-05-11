package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *config) setNext(nextUrl string) {
	c.Next = nextUrl
}
func (c *config) setPrevious(previousUrl string) {
	c.Previous = previousUrl
}

// for the things that both map commands have in common.
func mapCommandCommon(resourceURL string, currentConfig *config) error {
	// GET some information
	res, err := http.Get(resourceURL)
	if err != nil {
		return fmt.Errorf("ERROR is: %w", err)
	}
	defer res.Body.Close()
	//
	var locationAreas LocationAreas
	translateData := json.NewDecoder(res.Body)
	if err := translateData.Decode(&locationAreas); err != nil {
		return fmt.Errorf("ERROR is: %w", err)
	}
	// Update the currentConfig so it'll be accurate.
	currentConfig.setNext(locationAreas.Next)
	currentConfig.setPrevious(locationAreas.Previous)
	//
	allResults := []string{}
	for result := range locationAreas.Results {
		// fmt.Println(result, "is a result")
		// fmt.Println(locationAreas.Results[result].Name, "locationAreas.Results[result].Name")
		allResults = append(allResults, locationAreas.Results[result].Name)
	}

	// fmt.Println("locationAreas", locationAreas)
	for _, aResult := range allResults {
		fmt.Println(aResult)
	}
	return nil
}

func commandMap(currentConfig *config) error {
	// Reseting / keeping the currentConfig accurate.
	// fmt.Println(currentConfig.Next, "currentConfig.Next")
	// fmt.Println(strings.Contains(currentConfig.Next, "location-area"), "does it have 'location-area'? ")
	//
	if !strings.Contains(currentConfig.Next, "location-area") {
		currentConfig.setNext("https://pokeapi.co/api/v2/location-area")
		currentConfig.setPrevious("")
	}
	// fmt.Println(currentConfig.Next, "currentConfig.Next")
	//
	err := mapCommandCommon(currentConfig.Next, currentConfig)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}
	return nil

}

func commandMapB(currentConfig *config) error {
	if currentConfig.Previous != "" {
		currentConfig.setNext(currentConfig.Previous)
	}
	if currentConfig.Previous == "" {
		fmt.Println("you're on the first page")
		currentConfig.setNext("https://pokeapi.co/api/v2/location-area")
		return nil
	}
	//
	err := mapCommandCommon(currentConfig.Next, currentConfig)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}
	return nil
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
