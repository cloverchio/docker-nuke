package main

import (
	"fmt"
	"os"
	"github.com/cloverchio/docker-nuke/cmd"
)

func main() {
	subCommandError := cmd.ProcessSubcommandArgs(os.Args)
	if subCommandError != nil {
		fmt.Println(subCommandError)
		os.Exit(1)
	}
}
