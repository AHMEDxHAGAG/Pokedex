package commands

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var Commands = map[string]cliCommand{
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
