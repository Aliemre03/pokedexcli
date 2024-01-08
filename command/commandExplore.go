package command

import "fmt"

func CallbackExplore(config *Config, args ...string) error {
	if len(args) == 0 {
		return nil
	}

	locationAreaName := args[0]

	locationArea, err := config.PokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Location Area: %s\n", locationArea.Name)
	fmt.Printf("Encounter Method Rates:\n")
	for _, encounterMethodRate := range locationArea.EncounterMethodRates {
		fmt.Printf("\tEncounter Method: %s\n", encounterMethodRate.EncounterMethod.Name)
		for _, versionDetail := range encounterMethodRate.VersionDetails {
			fmt.Printf("\t\tVersion: %s\n", versionDetail.Version.Name)
			fmt.Printf("\t\tRate: %d\n", versionDetail.Rate)
		}
	}

	fmt.Printf("Pokemon Encounters:\n")
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Printf("\tPokemon: %s\n", pokemonEncounter.Pokemon.Name)
		for _, versionDetail := range pokemonEncounter.VersionDetails {
			fmt.Printf("\t\tVersion: %s\n", versionDetail.Version.Name)
			fmt.Printf("\t\tMax Chance: %d\n", versionDetail.MaxChance)
			for _, encounterDetail := range versionDetail.EncounterDetails {
				fmt.Printf("\t\t\tEncounter Method: %s\n", encounterDetail.Method.Name)
				fmt.Printf("\t\t\tMin Level: %d\n", encounterDetail.MinLevel)
				fmt.Printf("\t\t\tMax Level: %d\n", encounterDetail.MaxLevel)
				fmt.Printf("\t\t\tChance: %d\n", encounterDetail.Chance)
			}
		}
	}

	return nil
}
