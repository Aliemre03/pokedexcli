package command

import (
	"fmt"
	"os"
)

func CallbackExit(config *Config, args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
