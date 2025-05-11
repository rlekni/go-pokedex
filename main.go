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
		if len(cleanedInput) > 0 {
			fmt.Println("Your command was:", cleanedInput[0])
		}
	}
}
