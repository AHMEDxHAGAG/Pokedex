// Package commands
package commands

type Config struct {
	Reg registry
}

func NewConfig() *Config {
	return &Config{Reg: getRegistery()}
}
