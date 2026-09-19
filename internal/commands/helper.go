package commands

import (
	"io"
	"net/http"
)

func GetObject(key string, conf *Config) ([]byte, error) {
	val, found := conf.Cache.Get(key)
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
		conf.Cache.Add(key, val)
	}
	return val, nil
}
