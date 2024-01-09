package command

import (
	"bufio"
	"fmt"
	"github.com/Aliemre03/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type Config struct {
	PokeapiClient *pokeapi.Client
	nextLocation  *string
	prevLocation  *string
	CaughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints this help",
			callback:    CallbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    CallbackExit,
		},
		"map": {
			name:        "map",
			description: "List location areas",
			callback:    CallbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previoes location areas",
			callback:    CallbackMapb,
		},
		"explore": {
			name:        "explore <location-area-name>",
			description: "List the pokemon in the location area",
			callback:    CallbackExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch a pokemon in the location area",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the caught pokemon",
			callback:    callbackPokedex,
		},
	}
}

func StartRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := cleaned[1:]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(config, args...)
		if err != nil {
			return
		}

	}
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
