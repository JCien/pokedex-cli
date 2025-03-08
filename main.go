package main

import (
	"time"

	"github.com/JCien/pokedex-cli/internal/pokedexapi"
)

func main() {
	pokeClient := pokedexapi.NewClient(5 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
