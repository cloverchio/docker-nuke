package client

import (
	"fmt"
	"github.com/docker/docker/client"
)

var Docker *client.Client

func InitializeDockerClient() {
	dockerClient, dockerClientError := client.NewClientWithOpts(client.FromEnv)
	if dockerClientError != nil {
		fmt.Printf("Error initializing Docker client: %v", dockerClientError)
		return
	}
	Docker = dockerClient
}
