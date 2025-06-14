package client

import (
	"fmt"
	"github.com/docker/docker/client"
)

var Docker *client.Client

func InitializeDockerClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("Error initializing Docker client: %v", err)
		return
	}
	Docker = cli
}