package pokeapi

import (
	"io"
	"net/http"

	"github.com/AHMEDxHAGAG/Pokedex/internal/pokecache"
)

func GetObject(key string, cache *pokecache.Cache) ([]byte, error) {
	val, found := cache.Get(key)
	if !found {
		req, err := http.Get(key)
		if err != nil {
			return nil, err
		}
		defer func() { _ = req.Body.Close() }()
		val, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		cache.Add(key, val)
	}
	return val, nil
}
