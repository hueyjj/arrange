package main

import (
	"os"

	"github.com/hueyjj/arrange-server/cmd/arrange/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
