package main

import (
	"fmt"
	"os"
	"github.com/cloverchio/docker-nuke/cmd"
)

func main() {
	ok, err := cmd.ProcessSubcommandArgs(os.Args)
	if !ok {
		fmt.Println(err)
		os.Exit(1)
	}
}
