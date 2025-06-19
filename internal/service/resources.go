package service

import "github.com/cloverchio/docker-nuke/internal/client"

// RemoveAllResources removes all Docker resources, including containers, images,
// volumes, and networks.
//
// This function first calls RemoveAllContainers to remove all containers, then
// sequentially calls RemoveAllImages, RemoveAllVolumes, and RemoveAllNetworks
// to remove the corresponding resources. If any errors occur during the removal
// of any resource, the function will return the encountered error immediately.
// The operations are performed in sequence, and if any step fails, the process
// is halted and the error is returned.
//
// Parameters:
//   - dockerClient: A client instance that conforms to the Docker interface,
//     used for interacting with Docker containers, images, volumes, and networks.
//
// Returns:
//   - error: Returns an error if any of the resource removal functions fail.
//     If all operations succeed, it returns nil.
//
// Example Usage:
//   err := RemoveAllResources(dockerClient)
//   if err != nil {
//     fmt.Println("Error removing resources:", err)
//   }
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
