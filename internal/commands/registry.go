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
			Name:        "explore <location area>",
			Description: "List all Pokemons located in area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch <pokemon name>",
			Description: "Catch Pokemon and add it to the Pokedex",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect <pokemon name>",
			Description: "Inspect details about your catched pokemon",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all your catched pokemon",
			Callback:    commandPokedex,
		},
	}
}
