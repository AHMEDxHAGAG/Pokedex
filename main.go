package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputCleaned := cleanInput(input)
		fmt.Printf("Your command was: %s\n", inputCleaned[0])
	}
}
