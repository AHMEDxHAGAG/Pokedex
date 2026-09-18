// Package commands
package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const url = "https://pokeapi.co/api/v2/location-area/"

func commandExit(conf *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range conf.Reg {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}
	return nil
}

func commandMap(conf *Config) error {
	var location struct {
		Name string `json:"name"`
	}

	position, err := strconv.Atoi(conf.NextMap)
	if err != nil {
		return err
	}
	for i := position - 20; i < position; i++ {
		res, err := http.Get(fmt.Sprintf("%s%d", url, i))
		if err != nil {
			return err
		}
		defer func() { _ = res.Body.Close() }()
		if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
			return err
		}
		fmt.Println(location.Name)
	}
	conf.PrevMap = conf.NextMap
	conf.NextMap = fmt.Sprintf("%d", position+20)
	return nil
}

func commandMapb(conf *Config) error {
	var location struct {
		Name string `json:"name"`
	}

	position, err := strconv.Atoi(conf.PrevMap)
	if err != nil {
		return err
	}
	for i := position - 20; i < position; i++ {
		res, err := http.Get(fmt.Sprintf("%s%d", url, i))
		if err != nil {
			return err
		}
		defer func() { _ = res.Body.Close() }()
		if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
			return err
		}
		fmt.Println(location.Name)
	}
	conf.NextMap = conf.PrevMap
	conf.PrevMap = fmt.Sprintf("%d", position-20)
	return nil
}
