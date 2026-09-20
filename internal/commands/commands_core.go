// Package commands
package commands

import (
	"fmt"
	"os"
)

func commandExit(conf *Config, parameters []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config, parameters []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("\tUsage:")
	for _, value := range conf.Reg {
		fmt.Printf("\t\t%s: %s\n", value.Name, value.Description)
	}
	return nil
}
