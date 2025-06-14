package cmd

import (
	"errors"
	"fmt"
	"github.com/cloverchio/docker-nuke/internal/flags"
	"github.com/cloverchio/docker-nuke/pkg"
)

func ProcessNuke(subCommands []string) (bool, error) {
	nuke := flags.NukeFlagSet()
	nuke.Usage = pkg.Usage()
	err := nuke.Parse(subCommands)
	if err != nil {
		nuke.Usage()
		return false, fmt.Errorf("Error parsing flags: %v", err)
	}
	if *flags.All {
		fmt.Println(pkg.SubCommandMessage("unused", "containers, images, volumes, and networks"))
		return true, nil
	}
	if *flags.Containers {
		fmt.Println(pkg.SubCommandMessage("stopped", "containers"))
	}
	if *flags.Images {
		fmt.Println(pkg.SubCommandMessage("dangling", "images"))
	}
	if *flags.AllImages {
		fmt.Println(pkg.SubCommandMessage("unused", "images"))
	}
	if *flags.Volumes {
		fmt.Println(pkg.SubCommandMessage("unused", "volumes"))
	}
	if *flags.Networks {
		fmt.Println(pkg.SubCommandMessage("unused", "networks"))
	}
	if !(*flags.Containers || *flags.Images || *flags.AllImages || *flags.Volumes || *flags.Networks) {
		nuke.Usage()
		return false, errors.New("Please specify at least one flag to nuke: --containers, --images, --volumes, --networks, or --all")
	}
	return true, nil
}
