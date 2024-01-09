package command

import "fmt"

func callbackPokedex(config *Config, args ...string) error {
	fmt.Printf("You have %d pokemon\n", len(config.CaughtPokemon))

	for _, pokemon := range config.CaughtPokemon {
		fmt.Printf("%s\n", pokemon.Name)
	}
	return nil
}
