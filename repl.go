package main

import (
	"strings"
)

func ClearInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

// var commands =
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
