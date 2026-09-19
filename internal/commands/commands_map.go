// Package commands
package commands

import (
	"encoding/json"
	"fmt"
	"io"
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

/*
	TODO:
	Update your code that makes requests to the PokeAPI to use the cache.
	Create the cache once and reuse it in your PokeAPI request layer.
	If you already have the data for a given URL (which is our cache key) in the cache,
	you should use that instead of making a new request.
	Whenever you do make a request,
	you should add the response to the cache.
*/

func commandMap(conf *Config) error {
	if conf.NextMap == "null" {
		fmt.Println("No More Location Areas")
		return nil
	}
	val, found := conf.Cache.Get(conf.NextMap)
	if !found {
		req, err := http.Get(conf.NextMap)
		if err != nil {
			return err
		}
		defer func() { _ = req.Body.Close() }()
		val, err = io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		conf.Cache.Add(conf.NextMap, val)
	}
	if err := json.Unmarshal(val, &locationArea); err != nil {
		return err
	}
	conf.PrevMap = conf.NextMap
	conf.NextMap = locationArea.Next
	results := locationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb(conf *Config) error {
	if conf.PrevMap == "null" {
		fmt.Println("There is No Previous Location Areas")
		return nil
	}
	val, found := conf.Cache.Get(conf.PrevMap)
	if !found {
		req, err := http.Get(conf.PrevMap)
		if err != nil {
			return err
		}
		defer func() { _ = req.Body.Close() }()
		val, err = io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		conf.Cache.Add(conf.PrevMap, val)
	}
	if err := json.Unmarshal(val, &locationArea); err != nil {
		return err
	}
	conf.NextMap = conf.PrevMap
	conf.PrevMap = locationArea.Prev
	results := locationArea.Results
	for _, val := range results {
		fmt.Println(val.Name)
	}
	return nil
}
