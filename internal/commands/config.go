package commands

import (
	"time"

	"github.com/AHMEDxHAGAG/Pokedex/internal/pokeapi"
	"github.com/AHMEDxHAGAG/Pokedex/internal/pokecache"
)

type Config struct {
	Reg     Registry
	NextMap string
	PrevMap string
	Cache   *pokecache.Cache
	Pokedex map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	return &Config{
		Reg:     GetRegistry(),
		NextMap: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		PrevMap: "",
		Cache:   pokecache.NewCache(10 * time.Minute),
		Pokedex: map[string]pokeapi.Pokemon{},
	}
}
