package commands

type registry map[string]cliCommand

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func getRegistery() registry {
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
	}
}
