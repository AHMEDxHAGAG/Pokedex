// Package commands
package commands

import (
	"fmt"
	"os"
)

func commandExit(conf *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range conf.Reg {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}
	return nil
}
