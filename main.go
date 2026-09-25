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
		input_text := scanner.Text()
		clean_text := cleanInput(input_text)
		first_word := clean_text[0]
		fmt.Printf("Your command was:%s\n", first_word)
		scanner.Err()
	}
}
