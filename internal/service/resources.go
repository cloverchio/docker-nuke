package service

import "github.com/cloverchio/docker-nuke/internal/client"

func RemoveAllResources(dockerClient client.Docker) error {
    containerRemoveError := RemoveAllContainers(dockerClient)
    if containerRemoveError != nil {
        return containerRemoveError
    }
    imageRemoveError := RemoveAllImages(dockerClient)
    if imageRemoveError != nil {
        return imageRemoveError
    }
    volumeRemoveError := RemoveAllVolumes(dockerClient)
    if volumeRemoveError != nil {
        return volumeRemoveError
    }
    networkRemoveError := RemoveAllNetworks(dockerClient)
    if networkRemoveError != nil {
        return networkRemoveError
    }
    return nil
}
