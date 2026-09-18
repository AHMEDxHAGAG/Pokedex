package main

import "strings"

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
