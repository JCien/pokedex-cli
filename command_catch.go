package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide a Pokemon name.")
	}

	name := args[0]
	pokemonInfo, err := cfg.pokeapiClient.PokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])

	catchVar := int(pokemonInfo.BaseExperience / 10)

	if rand.Intn(catchVar) == int(catchVar/2) {
		fmt.Printf("%v was caught!\n", name)
		cfg.caughtPokemon[pokemonInfo.Name] = pokemonInfo
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}
