// Package commands
package commands

import (
	"github.com/AHMEDxHAGAG/Pokedex/internal/pokecache"
)

type Config struct {
	Reg     registry
	NextMap string
	PrevMap string
	Cache   *pokecache.Cache
}

func NewConfig() *Config {
	return &Config{
		Reg:     getRegistry(),
		NextMap: "https://pokeapi.co/api/v2/location-area/",
		PrevMap: "null",
		Cache:   pokecache.NewCache(5),
	}
}
