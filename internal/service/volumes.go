package service

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/volume"
)

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
