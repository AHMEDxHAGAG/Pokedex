package commands

type Registry map[string]cliCommand

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

func GetRegistry() Registry {
	return Registry{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of the next 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of the previous 20 location areas in the Pokemon world",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "List all Pokemons located in area : explore <location area>",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch Pokemon and add it to the Pokedex : catch <pokemon name>",
			Callback:    commandCatch,
		},
	}
}
