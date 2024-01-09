package main

import (
	"github.com/Aliemre03/pokedexcli/command"
	"github.com/Aliemre03/pokedexcli/internal/pokeapi"
)

func main() {
	config := command.Config{
		PokeapiClient: pokeapi.NewClient(),
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	command.CallbackHelp(&config)
	command.StartRepl(&config)
}
