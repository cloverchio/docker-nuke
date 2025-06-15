package main

import (
	"fmt"
	"os"
	"github.com/cloverchio/docker-nuke/cmd"
	"github.com/cloverchio/docker-nuke/internal/client"
)

func main() {
	client.InitializeDockerClient()
	subCommandError := cmd.ProcessSubcommandArgs(os.Args)
	if subCommandError != nil {
		fmt.Println(subCommandError)
		os.Exit(1)
	}
}
