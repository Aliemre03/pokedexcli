package command

import (
	"fmt"
	"os"
)

func CallbackExit(config *Config) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
