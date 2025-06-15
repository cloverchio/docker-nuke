package service

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/container"
)

func StopAllContainers() error {
    containers, containerListError := client.Docker.ContainerList(context.Background(), container.ListOptions{All: true})
    if containerListError != nil {
        fmt.Printf("Error listing containers: %v\n", containerListError)
        return containerListError
    }
    for _, individualContainer := range containers {
        containerStopError := client.Docker.ContainerStop(context.Background(), individualContainer.ID, container.StopOptions{})
        if containerStopError != nil {
            fmt.Printf("Error stopping container %s: %v\n", individualContainer.ID, containerStopError)
            return containerStopError
        }
        fmt.Printf("Successfully stopped container %s\n", individualContainer.ID)
    }
    return nil
}

func RemoveAllContainers() error {
    containerStopError := StopAllContainers()
    if containerStopError != nil {
        return containerStopError
    }
    containers, containerListError := client.Docker.ContainerList(context.Background(), container.ListOptions{All: true})
    if containerListError != nil {
        containerListErrorMessage(containerListError)
        return containerListError
    }
    for _, individualContainer := range containers {
        containerRemoveError := client.Docker.ContainerRemove(context.Background(), individualContainer.ID, container.RemoveOptions{Force: true})
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
