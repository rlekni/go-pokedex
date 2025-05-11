package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func (c cliCommand) String() string {
	return fmt.Sprintf("\t%s: %s", c.name, c.description)
}

const helpText = `
	Welcome to the Pokedex!
	Usage:
`

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(helpText)
	for _, v := range getCommands() {
		fmt.Println(v.String())
	}
	return nil
}
