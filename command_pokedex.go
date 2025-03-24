package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("No Pokemon caught yet")
		return nil
	}

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("- " + pokemon.Name)
	}

	return nil
}
