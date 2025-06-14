package object

import (
    "context"
    "fmt"
    "github.com/cloverchio/docker-nuke/internal/client"
    "github.com/docker/docker/api/types/filters"
)

func RemoveAllVolumes() error {
    containerStopError := StopAllContainers()
    if containerStopError != nil {
        return containerStopError
    }
    volumesList, volumeListError := client.Docker.VolumeList(context.Background(), filters.Args{})
    if volumeListError != nil {
        fmt.Printf("Error listing volumes: %v\n", volumeListError)
        return volumeListError
    }
    for _, volume := range volumesList.Volumes {
        volumeRemoveError := client.Docker.VolumeRemove(context.Background(), volume.Name, true)
        if volumeRemoveError != nil {
            fmt.Printf("Error removing volume %s: %v\n", volume.Name, volumeRemoveError)
            return volumeRemoveError
        }
        fmt.Printf("Successfully removed volume %s\n", volume.Name)
    }
    return nil
}
