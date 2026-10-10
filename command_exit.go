package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, param string) error {
	if param != "" {
		return errors.New("exit command accepts no parameters")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
