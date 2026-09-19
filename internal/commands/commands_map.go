// Package commands
package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	if conf.NextMap == "null" {
		fmt.Println("No More Location Areas")
		return nil
	}
	req, err := http.Get(conf.NextMap)
	if err != nil {
		return err
	}
	defer func() { _ = req.Body.Close() }()
	if err := json.NewDecoder(req.Body).Decode(&locationArea); err != nil {
		return err
	}
	conf.NextMap = locationArea.Next
	conf.PrevMap = locationArea.Prev
	results := locationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb(conf *Config) error {
	if conf.NextMap == "null" {
		fmt.Println("There is No Previous Location Areas")
		return nil
	}
	req, err := http.Get(conf.PrevMap)
	if err != nil {
		return err
	}
	defer func() { _ = req.Body.Close() }()
	if err := json.NewDecoder(req.Body).Decode(&locationArea); err != nil {
		return err
	}
	conf.NextMap = locationArea.Next
	conf.PrevMap = locationArea.Prev
	results := locationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}
