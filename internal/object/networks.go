package object

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/cloverchio/docker-nuke/internal/client"
)

func RemoveAllNetworks() error {
	containerStopError := StopAllContainers()
	if containerStopError != nil {
		return containerStopError
	}
	networks, networkListError := client.Docker.NetworkList(context.Background(), types.NetworkListOptions{})
	if networkListError != nil {
		fmt.Printf("Error listing networks: %v\n", networkListError)
		return networkListError
	}
	for _, network := range networks {
		networkRemoveError := client.Docker.NetworkRemove(context.Background(), network.ID)
		if networkRemoveError != nil {
			fmt.Printf("Error removing network %s: %v\n", network.ID, networkRemoveError)
			return networkRemoveError
		}
		fmt.Printf("Successfully removed network %s\n", network.ID)
	}
	return nil
}
