// Package commands
package commands

import (
	"encoding/json"
	"fmt"
)

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var locationArea struct {
	Next    string   `json:"next"`
	Prev    string   `json:"previous"`
	Results []result `json:"results"`
}

func commandMap(conf *Config) error {
	if conf.NextMap == "" {
		fmt.Println("No More Location Areas")
		return nil
	}

	val, err := GetObject(conf.NextMap, conf)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, &locationArea); err != nil {
		return err
	}

	conf.PrevMap = locationArea.Prev
	conf.NextMap = locationArea.Next
	results := locationArea.Results

	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb(conf *Config) error {
	if conf.PrevMap == "" {
		fmt.Println("There is No Previous Location Areas")
		return nil
	}

	val, err := GetObject(conf.PrevMap, conf)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, &locationArea); err != nil {
		return err
	}

	conf.PrevMap = locationArea.Prev
	conf.NextMap = locationArea.Next
	results := locationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}
