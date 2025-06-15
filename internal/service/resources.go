package service

func RemoveAllResources() error {
    containerRemoveError := RemoveAllContainers()
    if containerRemoveError != nil {
        return containerRemoveError
    }
    imageRemoveError := RemoveAllImages()
    if imageRemoveError != nil {
        return imageRemoveError
    }
    volumeRemoveError := RemoveAllVolumes()
    if volumeRemoveError != nil {
        return volumeRemoveError
    }
    networkRemoveError := RemoveAllNetworks()
    if networkRemoveError != nil {
        return networkRemoveError
    }
    return nil
}
