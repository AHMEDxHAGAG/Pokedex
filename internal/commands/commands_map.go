// Package commands
package commands

import (
	"encoding/json/v2"
	"fmt"

	"github.com/AHMEDxHAGAG/Pokedex/internal/pokemonapi"
)

func commandMap(conf *Config, parameters []string) error {
	if conf.NextMap == "" {
		fmt.Println("No More Location Areas")
		return nil
	}

	val, err := GetObject(conf.NextMap, conf)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, &pokemonapi.LocationArea); err != nil {
		return err
	}

	conf.PrevMap = pokemonapi.LocationArea.Prev
	conf.NextMap = pokemonapi.LocationArea.Next
	results := pokemonapi.LocationArea.Results

	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb(conf *Config, parameters []string) error {
	if conf.PrevMap == "" {
		fmt.Println("There is No Previous Location Areas")
		return nil
	}

	val, err := GetObject(conf.PrevMap, conf)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, &pokemonapi.LocationArea); err != nil {
		return err
	}

	conf.PrevMap = pokemonapi.LocationArea.Prev
	conf.NextMap = pokemonapi.LocationArea.Next
	results := pokemonapi.LocationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}
