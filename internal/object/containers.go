package object

import (
    "context"
    "fmt"
    "github.com/docker/docker/api/types"
    "github.com/cloverchio/docker-nuke/internal/client"
)

func StopAllContainers() error {
    containers, containerListError := client.Docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})
    if containerListError != nil {
        fmt.Printf("Error listing containers: %v\n", containerListError)
        return containerListError
    }
    for _, container := range containers {
        containerStopError := client.Docker.ContainerStop(context.Background(), container.ID, nil)
        if containerStopError != nil {
            fmt.Printf("Error stopping container %s: %v\n", container.ID, containerStopError)
            return containerStopError
        }
        fmt.Printf("Successfully stopped container %s\n", container.ID)
    }
    return nil
}

func RemoveAllContainers() error {
    containerStopError := StopAllContainers()
    if containerStopError != nil {
        return containerStopError
    }
    containers, containerListError := client.Docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})
    if containerListError != nil {
        containerListErrorMessage(containerListError)
        return containerListError
    }
    for _, container := range containers {
        containerRemoveError := client.Docker.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{Force: true})
        if containerRemoveError != nil {
            fmt.Printf("Error removing container %s: %v\n", container.ID, containerRemoveError)
            return containerRemoveError
        }
        fmt.Printf("Successfully removed container %s\n", container.ID)
    }
    return nil
}

func containerListErrorMessage(containerListError error) {
    fmt.Printf("Error listing containers: %v\n", containerListError)
}
