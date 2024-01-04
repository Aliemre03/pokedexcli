package command

import (
	"fmt"
)

func CallbackHelp() error {
	fmt.Println("Available commands:")
	for _, command := range getCommands() {
		fmt.Printf("  %s: %s\n", command.name, command.description)

	}

	return nil
}

//fmt.Printf("  %s: %s\n", command.name, command.description)
