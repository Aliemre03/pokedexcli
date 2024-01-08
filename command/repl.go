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
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(config)

	}
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
