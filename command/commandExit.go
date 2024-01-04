package command

import (
	"fmt"
	"os"
)

func CallbackExit() error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
