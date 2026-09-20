package commands

import (
	"encoding/json/v2"
	"fmt"
	"math/rand"

	"github.com/AHMEDxHAGAG/Pokedex/internal/pokeapi"
)

func commandExplore(conf *Config, parameters []string) error {
	if len(parameters) < 1 {
		return fmt.Errorf("expected no. of arguments: %d, found: %d", 2, len(parameters))
	}
	locationArea := parameters[1]
	fmt.Println("Exploring " + locationArea)
	url := "https://pokeapi.co/api/v2/location-area/"
	api := url + locationArea
	val, err := GetObject(api, conf)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(val, &pokeapi.Location); err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, ecounters := range pokeapi.Location.PokemonEncounters {
		fmt.Printf("- %s\n", ecounters.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *Config, parameters []string) error {
	if len(parameters) < 1 {
		return fmt.Errorf("expected no. of arguments: %d, found: %d", 2, len(parameters))
	}
	pokeName := parameters[1]
	fmt.Println("Throwing a Pokeball at " + pokeName + "...")
	url := "https://pokeapi.co/api/v2/pokemon/"
	api := url + pokeName
	val, err := GetObject(api, conf)
	if err != nil {
		return err
	}
	var pokemonInstance pokeapi.Pokemon
	if err := json.Unmarshal(val, &pokemonInstance); err != nil {
		return err
	}
	userChance := rand.Intn(pokemonInstance.BaseExperience)
	caught := userChance <= 40
	if caught {
		fmt.Println(pokeName + " was caught!")
		conf.Pokedex[pokeName] = pokemonInstance
	} else {
		fmt.Printf("%s escaped!\n", pokeName)
	}
	return nil
}

func commandInspect(conf *Config, parameters []string) error {
	if len(parameters) < 1 {
		return fmt.Errorf("expected no. of arguments: %d, found: %d", 2, len(parameters))
	}
	pokeName := parameters[1]
	poke, caught := conf.Pokedex[pokeName]
	if !caught {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}
	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Printf("Stats:\n")
	for _, val := range poke.Stats {
		fmt.Printf("\t-%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, val := range poke.Types {
		fmt.Printf("\t- %s\n", val.Type.Name)
	}
	return nil
}

func commandPokedex(conf *Config, parameters []string) error {
	pokedex := conf.Pokedex
	if len(pokedex) == 0 {
		fmt.Println("you have not caught any pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf("\t- %s\n", name)
	}
	return nil
}
