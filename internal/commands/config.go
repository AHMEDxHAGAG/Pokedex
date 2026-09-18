// Package commands
package commands

type Config struct {
	Reg     registry
	NextMap string
	PrevMap string
}

func NewConfig() *Config {
	return &Config{Reg: getRegistry(), NextMap: "21", PrevMap: "1"}
}
