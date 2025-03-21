package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide a valid Pokemon name.")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if exists {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")

		for stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", pokemon.Stats[stat].Stat.Name, pokemon.Stats[stat].BaseStat)
		}

		fmt.Println("Types:")
		for pokemon_type := range pokemon.Types {
			fmt.Printf("  - %v\n", pokemon.Types[pokemon_type].Type.Name)
		}

		return nil
	} else {
		fmt.Println("You have not caught that Pokemon.")
		return nil
	}
}
