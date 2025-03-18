package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide a location name.")
	}

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _, poke := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %v\n", poke.Pokemon.Name)
	}

	return nil
}
