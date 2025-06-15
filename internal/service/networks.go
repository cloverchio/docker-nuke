package service

import (
	"context"
	"fmt"
	"github.com/cloverchio/docker-nuke/internal/client"
	"github.com/docker/docker/api/types/network"
	"slices"
)

func RemoveAllNetworks() error {
	containerStopError := StopAllContainers()
	if containerStopError != nil {
		return containerStopError
	}
	networks, networkListError := client.Docker.NetworkList(context.Background(), network.ListOptions{})
	if networkListError != nil {
		fmt.Printf("Error listing networks: %v\n", networkListError)
		return networkListError
	}
	defaultNetworks := []string{"bridge", "none", "host"}
	for _, individualNetwork := range networks {
		if slices.Contains(defaultNetworks, individualNetwork.Name) {
			continue
		}
		networkRemoveError := client.Docker.NetworkRemove(context.Background(), individualNetwork.ID)
		if networkRemoveError != nil {
			fmt.Printf("Error removing network %s: %v\n", individualNetwork.ID, networkRemoveError)
			return networkRemoveError
		}
		fmt.Printf("Successfully removed network %s\n", individualNetwork.ID)
	}
	return nil
}
