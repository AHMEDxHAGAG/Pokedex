package commands

import (
	"encoding/json/v2"
	"fmt"
	"math/rand"
)

var location struct {
	PokemonEncounters []pokemonEncounters `json:"pokemon_encounters"`
}

type pokemonEncounters struct {
	Pokemon namedAPIResource `json:"pokemon"`
}

type Pokemon struct {
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	BaseExperience int     `json:"base_experience"`
	Stats          []stats `json:"stats"`
	Types          []types `json:"types"`
}

type stats struct {
	Stat     namedAPIResource `json:"stat"`
	BaseStat int              `json:"base_stat"`
}

type types struct {
	Type namedAPIResource `json:"type"`
}

type namedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

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
	if err := json.Unmarshal(val, &location); err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, ecounters := range location.PokemonEncounters {
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
	var pokemonInstance Pokemon
	if err := json.Unmarshal(val, &pokemonInstance); err != nil {
		return err
	}
	userChance := rand.Intn(pokemonInstance.BaseExperience)
	caught := userChance <= 40
	if caught {
		fmt.Println(pokeName + " was caught!")
		conf.Pokedex[pokeName] = pokemonInstance
	} else {
		return fmt.Errorf(pokeName + " escaped!")
	}
	return nil
}

func commandInspect(conf *Config, parameters []string) error {
	if len(parameters) < 1 {
		return fmt.Errorf("expected no. of arguments: %d, found: %d", 2, len(parameters))
	}
	pokeName := parameters[1]
	_, caught := conf.Pokedex[pokeName]
	if !caught {
		return fmt.Errorf("you have not caught that pokemon")
	}
	return nil
}
