package service

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/image"
)

// RemoveDanglingImages removes all dangling Docker images.
//
// This function first stops all containers by calling StopAllContainers,
// then retrieves a list of all Docker images. It iterates over the images
// and removes those that do not have any repository tags (dangling images).
// If any errors occur during stopping containers, listing images, or removing
// images, the function will return the encountered error.
//
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker images and containers.
//
// Returns:
//   - error: Returns an error if there is an issue with stopping containers,
//     listing images, or removing dangling images. If all operations succeed,
//     it returns nil.
//
// Example Usage:
//   err := RemoveDanglingImages(dockerClient)
//   if err != nil {
//     fmt.Println("Error removing dangling images:", err)
//   }
func RemoveDanglingImages(dockerClient client.Docker) error {
    containerStopError := StopAllContainers(dockerClient)
    if containerStopError != nil {
        return containerStopError
    }
    images, imageListError := dockerClient.ImageList(context.Background(), image.ListOptions{All: true})
    if imageListError != nil {
        imageListErrorMessage(imageListError)
        return imageListError
    }
    for _, individualImage := range images {
        if len(individualImage.RepoTags) == 0 {
            _, imageRemoveError := dockerClient.ImageRemove(context.Background(), individualImage.ID, image.RemoveOptions{Force: true})
            if imageRemoveError != nil {
                imageRemoveErrorMessage(individualImage.ID, imageListError)
                return imageRemoveError
            }
        }
    }
    return nil
}

// RemoveAllImages removes all Docker images.
//
// This function first stops all containers by calling StopAllContainers,
// then retrieves a list of all Docker images. It iterates over the images
// and removes each one. If any errors occur during stopping containers,
// listing images, or removing images, the function will return the encountered error.
//
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker images and containers.
//
// Returns:
//   - error: Returns an error if there is an issue with stopping containers,
//     listing images, or removing images. If all operations succeed, it returns nil.
//
// Example Usage:
//   err := RemoveAllImages(dockerClient)
//   if err != nil {
//     fmt.Println("Error removing all images:", err)
//   }
func RemoveAllImages(dockerClient client.Docker) error {
    containerStopError := StopAllContainers(dockerClient)
    if containerStopError != nil {
        return containerStopError
    }
    images, imageListError := dockerClient.ImageList(context.Background(), image.ListOptions{All: true})
    if imageListError != nil {
        imageListErrorMessage(imageListError)
        return imageListError
    }
    for _, individualImage := range images {
        _, imageRemoveError := dockerClient.ImageRemove(context.Background(), individualImage.ID, image.RemoveOptions{Force: true})
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
