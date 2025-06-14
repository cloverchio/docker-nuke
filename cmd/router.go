package cmd

import (
	"errors"
	"fmt"
)

func ProcessSubcommandArgs(args []string) (bool, error) {
	if len(args) < 2 {
		return false, errors.New("Must specifiy a sub-command: nuke or help")
	}
	switch args[1] {
	case "nuke":
		return ProcessNuke(args[2:])
	case "help":
		return ProcessHelp()
	default:
		return false, fmt.Errorf("Unknown subcommand: %s\nUse 'docker-nuke help' for instructions", args[1])
	}
}
