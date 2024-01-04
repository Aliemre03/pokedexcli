package command

import (
	"fmt"
	"github.com/Aliemre03/pokedexcli/internal/pokeapi"
)

func CallbackMap() error {
	pokeapiClient := pokeapi.NewClient()
	location, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		return err
	}
	fmt.Printf("Location areas:\n")

	for number, locationArea := range location.Results {
		fmt.Printf("%d - %s\n", number, locationArea.Name)
	}

	return nil
}
