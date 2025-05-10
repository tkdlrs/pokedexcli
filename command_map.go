package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(currentConfig *config) error {
	res, err := http.Get(currentConfig.Next)
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

func commandMapB(ptr *config) error {
	fmt.Println("GO BACK. ")
	return nil
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
