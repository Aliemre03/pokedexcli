package command

import (
	"fmt"
	"math/rand"
)

func callbackCatch(config *Config, args ...string) error {
	if len(args) == 0 {
		return nil
	}

	pokemonName := args[0]

	pokemon, err := config.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	randomNumber := rand.Intn(pokemon.BaseExperience)
	if randomNumber > 66 {
		return fmt.Errorf("pokemon %s escaped", pokemon.Name)
	}

	fmt.Printf("pokemon %s catched\n", pokemon.Name)
	config.CaughtPokemon[pokemon.Name] = *pokemon
	return nil
}
