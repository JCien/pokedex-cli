package main

import (
	"fmt"
)

func commandExplore(cfg *config, area_name string) error {
	location_url := *cfg.nextLocationsURL + area_name
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&location_url)
	if err != nil {
		return err
	}

	for _, poke := range pokemonResp.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}

	return nil
}
