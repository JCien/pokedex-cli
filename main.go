package main

import (
	"time"

	"github.com/JCien/pokedex-cli/internal/pokedexapi"
)

func main() {
	pokeClient := pokedexapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config{
		caughtPokemon: map[string]pokedexapi.PokemonInfo{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
