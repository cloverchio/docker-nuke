package service

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/image"
)

func RemoveDanglingImages() error {
    containerStopError := StopAllContainers()
    if containerStopError != nil {
        return containerStopError
    }
    images, imageListError := client.Docker.ImageList(context.Background(), image.ListOptions{All: true})
    if imageListError != nil {
        imageListErrorMessage(imageListError)
        return imageListError
    }
    for _, individualImage := range images {
        if len(individualImage.RepoTags) == 0 {
            _, imageRemoveError := client.Docker.ImageRemove(context.Background(), individualImage.ID, image.RemoveOptions{Force: true})
            if imageRemoveError != nil {
                imageRemoveErrorMessage(individualImage.ID, imageListError)
                return imageRemoveError
            }
        }
    }
    return nil
}

func RemoveAllImages() error {
    containerStopError := StopAllContainers()
    if containerStopError != nil {
        return containerStopError
    }
    images, imageListError := client.Docker.ImageList(context.Background(), image.ListOptions{All: true})
    if imageListError != nil {
        imageListErrorMessage(imageListError)
        return imageListError
    }
    for _, individualImage := range images {
        _, imageRemoveError := client.Docker.ImageRemove(context.Background(), individualImage.ID, image.RemoveOptions{Force: true})
        if imageRemoveError != nil {
            imageRemoveErrorMessage(individualImage.ID, imageRemoveError)
            return imageRemoveError
        }
        imageRemoveSuccessMessage(individualImage.ID)
    }
    return nil
}

func imageListErrorMessage(imageListError error) {
    fmt.Printf("Error listing images: %v\n", imageListError)
}

func imageRemoveErrorMessage(id string, imageRemoveError error) {
    fmt.Printf("Error removing image %s: %v\n", id, imageRemoveError)
}

func imageRemoveSuccessMessage(id string) {
    fmt.Printf("Successfully removed image %s\n", id)
}
