package service

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/volume"
)

// RemoveAllVolumes removes all Docker volumes.
//
// This function first stops all containers by calling StopAllContainers,
// then retrieves a list of all Docker volumes. It iterates over the volumes
// and removes each one. If any errors occur during stopping containers,
// listing volumes, or removing volumes, the function will print the error
// and return it.
//
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker volumes and containers.
//
// Returns:
//   - error: Returns an error if there is an issue with stopping containers,
//     listing volumes, or removing volumes. If all operations succeed, it returns nil.
//
// Example Usage:
//   err := RemoveAllVolumes(dockerClient)
//   if err != nil {
//     fmt.Println("Error removing volumes:", err)
//   }
func RemoveAllVolumes(dockerClient client.Docker) error {
    containerStopError := StopAllContainers(dockerClient)
    if containerStopError != nil {
        return containerStopError
    }
    volumes, volumeListError := dockerClient.VolumeList(context.Background(), volume.ListOptions{})
    if volumeListError != nil {
        fmt.Printf("Error listing volumes: %v\n", volumeListError)
        return volumeListError
    }
    for _, individualVolume := range volumes.Volumes {
        volumeRemoveError := dockerClient.VolumeRemove(context.Background(), individualVolume.Name, true)
        if volumeRemoveError != nil {
            fmt.Printf("Error removing volume %s: %v\n", individualVolume.Name, volumeRemoveError)
            return volumeRemoveError
        }
        fmt.Printf("Successfully removed volume %s\n", individualVolume.Name)
    }
    return nil
}
