package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rlekni/go-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemons: map[string]pokeapi.Pokemon{},
		pokeapiClient:  pokeClient,
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := ClearInput(scanner.Text())
		command, ok := getCommands()[words[0]]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
