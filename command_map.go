package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.NextURL = locations.Next
	cfg.PreviousURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.PreviousURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.PreviousURL)
	if err != nil {
		return err
	}

	cfg.NextURL = locations.Next
	cfg.PreviousURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}