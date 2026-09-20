// Package commands
package commands

import (
	"encoding/json/v2"
	"fmt"

	"github.com/AHMEDxHAGAG/Pokedex/internal/pokeapi"
)

func commandMap(conf *Config, parameters []string) error {
	if conf.NextMap == "" {
		fmt.Println("No More Location Areas")
		return nil
	}

	val, err := pokeapi.GetObject(conf.NextMap, conf.Cache)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, &pokeapi.LocationArea); err != nil {
		return err
	}

	conf.PrevMap = pokeapi.LocationArea.Prev
	conf.NextMap = pokeapi.LocationArea.Next
	results := pokeapi.LocationArea.Results

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

	val, err := pokeapi.GetObject(conf.PrevMap, conf.Cache)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, &pokeapi.LocationArea); err != nil {
		return err
	}

	conf.PrevMap = pokeapi.LocationArea.Prev
	conf.NextMap = pokeapi.LocationArea.Next
	results := pokeapi.LocationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}
