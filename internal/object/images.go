package object

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types"
)

func RemoveDanglingImages() error {
    containerStopError := StopAllContainers()
    if containerStopError != nil {
        return containerStopError
    }
    images, imageListError := client.Docker.ImageList(context.Background(), types.ImageListOptions{All: true})
    if imageListError != nil {
        imageListErrorMessage(imageListError)
        return imageListError
    }
    for _, image := range images {
        if len(image.RepoTags) == 0 {
            _, imageRemoveError := client.Docker.ImageRemove(context.Background(), image.ID, types.ImageRemoveOptions{Force: true})
            if imageRemoveError != nil {
                imageRemoveErrorMessage(image.ID, imageListError)
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
    images, imageListError := client.Docker.ImageList(context.Background(), types.ImageListOptions{All: true})
    if imageListError != nil {
        imageListErrorMessage(imageListError)
        return imageListError
    }
    for _, image := range images {
        _, imageRemoveError := client.Docker.ImageRemove(context.Background(), image.ID, types.ImageRemoveOptions{Force: true})
        if imageRemoveError != nil {
            imageRemoveErrorMessage(image.ID, imageRemoveError)
            return imageRemoveError
        }
        imageRemoveSuccessMessage(image.ID)
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
