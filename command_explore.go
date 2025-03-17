package main

import (
	"fmt"
	"net/url"
)

func commandExplore(cfg *config, area_name string) error {
	parsedURL, err := url.Parse(*cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	defaultURL := parsedURL.Scheme + "://" + parsedURL.Hostname() + parsedURL.Path
	location_url := defaultURL + "/" + area_name

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&location_url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", area_name)
	fmt.Println("Found Pokemon:")

	for _, poke := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %v\n", poke.Pokemon.Name)
	}

	return nil
}
