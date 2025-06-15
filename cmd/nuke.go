package cmd

import (
	"errors"
	"fmt"
	"github.com/cloverchio/docker-nuke/internal/flag"
	"github.com/cloverchio/docker-nuke/internal/object"
	"github.com/cloverchio/docker-nuke/pkg"
)

func ProcessNuke(subCommands []string) (bool, error) {
	nuke := flag.NukeFlagSet()
	nuke.Usage = pkg.Usage()
	err := nuke.Parse(subCommands)
	if err != nil {
		nuke.Usage()
		return false, fmt.Errorf("Error parsing flags: %v", err)
	}
	if *flag.All {
		fmt.Println(pkg.SubCommandMessage("unused", "containers, images, volumes, and networks"))
		return true, nil
	}
	if *flag.Containers {
		fmt.Println(pkg.SubCommandMessage("stopped", "containers"))
		containerRemoveError := object.RemoveAllContainers()
		if containerRemoveError != nil {
			return false, containerRemoveError
		}
	}
	if *flag.Images {
		fmt.Println(pkg.SubCommandMessage("dangling", "images"))
		imageRemoveError := object.RemoveDanglingImages()
		if imageRemoveError != nil {
			return false, imageRemoveError
		}
	}
	if *flag.AllImages {
		fmt.Println(pkg.SubCommandMessage("unused", "images"))
		imageRemoveError := object.RemoveAllImages()
		if imageRemoveError != nil {
			return false, imageRemoveError
		}
	}
	if *flag.Volumes {
		fmt.Println(pkg.SubCommandMessage("unused", "volumes"))
		volumeRemoveError := object.RemoveAllVolumes()
		if volumeRemoveError != nil {
			return false, volumeRemoveError
		}
	}
	if *flag.Networks {
		fmt.Println(pkg.SubCommandMessage("unused", "networks"))
		networkRemoveError := object.RemoveAllNetworks()
		if networkRemoveError != nil {
			return false, networkRemoveError
		}
	}
	if !(*flag.Containers || *flag.Images || *flag.AllImages || *flag.Volumes || *flag.Networks) {
		nuke.Usage()
		return false, errors.New("Please specify at least one flag to nuke: --containers, --images, --volumes, --networks, or --all")
	}
	return true, nil
}
