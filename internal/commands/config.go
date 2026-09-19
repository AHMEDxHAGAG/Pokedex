package commands

import (
	"time"

	"github.com/AHMEDxHAGAG/Pokedex/internal/pokecache"
)

type Config struct {
	Reg     Registry
	NextMap string
	PrevMap string
	Cache   *pokecache.Cache
}

func NewConfig() *Config {
	return &Config{
		Reg:     GetRegistry(),
		NextMap: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		PrevMap: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Cache:   pokecache.NewCache(10 * time.Minute),
	}
}
