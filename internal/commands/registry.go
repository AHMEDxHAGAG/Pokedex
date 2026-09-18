package commands

type registry map[string]cliCommand

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func getRegistry() registry {
	return registry{
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
	}
}
