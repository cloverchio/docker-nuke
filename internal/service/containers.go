package service

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/container"
)

// StopAllContainers stops all running Docker containers.
//
// This function retrieves a list of all containers (both running and stopped),
// then iterates over each container and attempts to stop it. If any errors
// occur during the listing of containers or stopping a container, they are
// printed to the console, and the function returns the error encountered.
// 
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker containers.
//
// Returns:
//   - error: If any error occurs during the process (listing containers or
//     stopping containers), it will return the error. Otherwise, it returns
//     nil when all containers are successfully stopped.
//
// Example Usage:
//   err := StopAllContainers(dockerClient)
//   if err != nil {
//     fmt.Println("Error stopping containers:", err)
//   }
func StopAllContainers(dockerClient client.Docker) error {
    containers, containerListError := dockerClient.ContainerList(context.Background(), container.ListOptions{All: true})
    if containerListError != nil {
        containerListErrorMessage(containerListError)
        return containerListError
    }
    for _, individualContainer := range containers {
        containerStopError := dockerClient.ContainerStop(context.Background(), individualContainer.ID, container.StopOptions{})
        if containerStopError != nil {
            fmt.Printf("Error stopping container %s: %v\n", individualContainer.ID, containerStopError)
            return containerStopError
        }
        fmt.Printf("Successfully stopped container %s\n", individualContainer.ID)
    }
    return nil
}

// RemoveAllContainers stops and then removes all Docker containers.
//
// This function first calls StopAllContainers to stop all containers, and
// if no errors are encountered, it proceeds to remove each container.
// If any errors occur during the stopping or removing of containers, the
// function prints the error and returns it.
//
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker containers.
//
// Returns:
//   - error: Returns an error if any issues occur during the stopping or
//     removal process. If both operations succeed, it returns nil.
//
// Example Usage:
//   err := RemoveAllContainers(dockerClient)
//   if err != nil {
//     fmt.Println("Error removing containers:", err)
//   }
func RemoveAllContainers(dockerClient client.Docker) error {
    containerStopError := StopAllContainers(dockerClient)
    if containerStopError != nil {
        return containerStopError
    }
    containers, containerListError := dockerClient.ContainerList(context.Background(), container.ListOptions{All: true})
    if containerListError != nil {
        containerListErrorMessage(containerListError)
        return containerListError
    }
    for _, individualContainer := range containers {
        fmt.Println(individualContainer.ID)
        containerRemoveError := dockerClient.ContainerRemove(context.Background(), individualContainer.ID, container.RemoveOptions{Force: true})
        if containerRemoveError != nil {
            fmt.Printf("Error removing container %s: %v\n", individualContainer.ID, containerRemoveError)
            return containerRemoveError
        }
        fmt.Printf("Successfully removed container %s\n", individualContainer.ID)
    }
    return nil
}

func containerListErrorMessage(containerListError error) {
    fmt.Printf("Error listing containers: %v\n", containerListError)
}
