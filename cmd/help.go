package cmd

import "fmt"

func ProcessHelp() error {
	fmt.Println("Usage: docker-nuke [subcommand] [options]")
	fmt.Println("Subcommands:")
	fmt.Println("  nuke         Nukes Docker containers, images, volumes, and networks")
	fmt.Println("  help         Show this help message")
	return nil
}
