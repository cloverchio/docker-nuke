package cmd

import (
	"errors"
	"fmt"
	"github.com/cloverchio/docker-nuke/internal/flag"
	"github.com/cloverchio/docker-nuke/internal/service"
	"github.com/cloverchio/docker-nuke/pkg"
	"github.com/docker/docker/client"
)

func ProcessNuke(subCommands []string) error {
	dockerClient, dockerClientError := client.NewClientWithOpts(client.FromEnv)
	if dockerClientError != nil {
		return fmt.Errorf("Error initializing Docker client: %v", dockerClientError)
	}
	nuke := flag.NukeFlagSet()
	nuke.Usage = pkg.Usage()
	parseError := nuke.Parse(subCommands)
	if parseError != nil {
		nuke.Usage()
		return fmt.Errorf("Error parsing flags: %v", parseError)
	}
	if *flag.All {
		fmt.Println(pkg.SubCommandMessage("unused", "containers, images, volumes, and networks"))
		resourceRemoveError := service.RemoveAllResources(dockerClient)
		if resourceRemoveError != nil {
			return resourceRemoveError
		}
		return nil
	}
	if *flag.Containers {
		fmt.Println(pkg.SubCommandMessage("stopped", "containers"))
		containerRemoveError := service.RemoveAllContainers(dockerClient)
		if containerRemoveError != nil {
			return containerRemoveError
		}
	}
	if *flag.Images {
		fmt.Println(pkg.SubCommandMessage("dangling", "images"))
		imageRemoveError := service.RemoveDanglingImages(dockerClient)
		if imageRemoveError != nil {
			return imageRemoveError
		}
	}
	if *flag.AllImages {
		fmt.Println(pkg.SubCommandMessage("unused", "images"))
		imageRemoveError := service.RemoveAllImages(dockerClient)
		if imageRemoveError != nil {
			return imageRemoveError
		}
	}
	if *flag.Volumes {
		fmt.Println(pkg.SubCommandMessage("unused", "volumes"))
		volumeRemoveError := service.RemoveAllVolumes(dockerClient)
		if volumeRemoveError != nil {
			return volumeRemoveError
		}
	}
	if *flag.Networks {
		fmt.Println(pkg.SubCommandMessage("unused", "networks"))
		networkRemoveError := service.RemoveAllNetworks(dockerClient)
		if networkRemoveError != nil {
			return networkRemoveError
		}
	}
	if !(*flag.Containers || *flag.Images || *flag.AllImages || *flag.Volumes || *flag.Networks) {
		nuke.Usage()
		return errors.New("Please specify at least one flag to nuke: --containers, --images, --volumes, --networks, or --all")
	}
	return nil
}
