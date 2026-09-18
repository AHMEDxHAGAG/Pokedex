package main

import (
	"github.com/AHMEDxHAGAG/Pokedex/internal/commands"
	"github.com/AHMEDxHAGAG/Pokedex/internal/repl"
)

func main() {
	conf := commands.NewConfig()
	repl.StartREPL(conf)
}
