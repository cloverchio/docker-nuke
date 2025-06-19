package service

import (
	"context"
	"fmt"
	"github.com/cloverchio/docker-nuke/internal/client"
	"github.com/docker/docker/api/types/network"
	"slices"
)

// RemoveAllNetworks removes all Docker networks except for the default networks.
//
// This function first stops all containers by calling StopAllContainers,
// then retrieves a list of all Docker networks. It iterates over the networks
// and removes any network that is not one of the default networks: "bridge", "none",
// or "host". If any errors occur during stopping containers, listing networks, or
// removing networks, the function will print the error and return it.
//
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker networks and containers.
//
// Returns:
//   - error: Returns an error if there is an issue with stopping containers,
//     listing networks, or removing networks. If all operations succeed, it returns nil.
//
// Example Usage:
//   err := RemoveAllNetworks(dockerClient)
//   if err != nil {
//     fmt.Println("Error removing networks:", err)
//   }
func RemoveAllNetworks(dockerClient client.Docker) error {
	containerStopError := StopAllContainers(dockerClient)
	if containerStopError != nil {
		return containerStopError
	}
	networks, networkListError := dockerClient.NetworkList(context.Background(), network.ListOptions{})
	if networkListError != nil {
		fmt.Printf("Error listing networks: %v\n", networkListError)
		return networkListError
	}
	defaultNetworks := []string{"bridge", "none", "host"}
	for _, individualNetwork := range networks {
		if slices.Contains(defaultNetworks, individualNetwork.Name) {
			continue
		}
		networkRemoveError := dockerClient.NetworkRemove(context.Background(), individualNetwork.ID)
		if networkRemoveError != nil {
			fmt.Printf("Error removing network %s: %v\n", individualNetwork.ID, networkRemoveError)
			return networkRemoveError
		}
		fmt.Printf("Successfully removed network %s\n", individualNetwork.ID)
	}
	return nil
}
