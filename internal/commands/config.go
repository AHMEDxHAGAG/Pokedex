// Package commands
package commands

type Config struct {
	Reg     registry
	NextMap string
	PrevMap string
}

func NewConfig() *Config {
	return &Config{Reg: getRegistry(), NextMap: "https://pokeapi.co/api/v2/location-area/", PrevMap: "null"}
}
