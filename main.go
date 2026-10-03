package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	// next should start of with the first result of the API call
	cfg := config{commands, "https://pokeapi.co/api/v2/location-area/", ""}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		// commands found in repl.go
		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(&cfg); err != nil {
			fmt.Println(err)
		}

		scanner.Err()
	}
}
