// Package repl
package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AHMEDxHAGAG/Pokedex/internal/commands"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputCleaned := cleanInput(input)
		val, ok := commands.Commands[inputCleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := val.Callback(); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	sliceOfTXT := strings.Split(text, " ")
	newSlice := []string{}
	for _, val := range sliceOfTXT {
		if val == "" {
			continue
		}
		newSlice = append(newSlice, val)
	}
	return newSlice
}
