package command

import (
	"fmt"
)

func CallbackMap(config *Config) error {
	location, err := config.PokeapiClient.ListLocationAreas(config.nextLocation)
	if err != nil {
		return err
	}
	fmt.Printf("Location areas:\n")

	for number, locationArea := range location.Results {
		fmt.Printf("%d - %s\n", number, locationArea.Name)
	}

	config.nextLocation = location.Next
	config.prevLocation = location.Previous

	return nil
}

func CallbackMapb(config *Config) error {
	location, err := config.PokeapiClient.ListLocationAreas(config.prevLocation)
	if err != nil {
		return err
	}
	fmt.Printf("Location areas:\n")

	for number, locationArea := range location.Results {
		fmt.Printf("%d - %s\n", number, locationArea.Name)
	}

	config.nextLocation = location.Next
	config.prevLocation = location.Previous

	return nil
}
