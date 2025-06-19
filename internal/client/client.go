package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
)

type Docker interface {
	ContainerList(context context.Context, options container.ListOptions) ([]types.Container, error)
	ContainerStop(context context.Context, containerID string, options container.StopOptions) error
	ContainerRemove(context context.Context, containerID string, options container.RemoveOptions) error
	ImageList(context context.Context, options image.ListOptions) ([]image.Summary, error)
	ImageRemove(context context.Context, imageID string, options image.RemoveOptions) ([]image.DeleteResponse, error)
	NetworkList(context context.Context, options network.ListOptions) ([]network.Summary, error)
	NetworkRemove(context context.Context, networkID string) error
	VolumeList(context context.Context, options volume.ListOptions) (volume.ListResponse, error)
	VolumeRemove(context context.Context, volumeID string, force bool) error
}
