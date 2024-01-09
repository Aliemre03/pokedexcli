package command

import "fmt"

func callbackInspect(config *Config, args ...string) error {
	if len(args) == 0 {
		return nil
	}

	pokemonName := args[0]

	pokemon, ok := config.CaughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon %s not found", pokemonName)
	}

	fmt.Printf("pokemon %s\n", pokemon.Name)
	fmt.Printf("base experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("height: %d\n", pokemon.Height)
	fmt.Printf("weight: %d\n", pokemon.Weight)
	fmt.Printf("types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("%s ", t.Type.Name)

	}
	fmt.Printf("\n")
	fmt.Printf("stats: ")
	for _, s := range pokemon.Stats {
		fmt.Printf("%s: %d ", s.Stat.Name, s.BaseStat)

	}
	fmt.Printf("\n")
	fmt.Printf("abilities: ")
	for _, a := range pokemon.Abilities {
		fmt.Printf("%s ", a.Ability.Name)

	}
	fmt.Printf("\n")
	return nil
}
