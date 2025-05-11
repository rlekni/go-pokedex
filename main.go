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
		cleanedInput := ClearInput(scanner.Text())
		command, ok := getCommands()[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
