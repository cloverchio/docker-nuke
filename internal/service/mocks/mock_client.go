package mocks

import (
    "context"
    "github.com/stretchr/testify/mock"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/api/types/image"
    "github.com/docker/docker/api/types/network"
    "github.com/docker/docker/api/types/volume"
)

type MockDockerClient struct {
    mock.Mock
}

func (mockDockerClient *MockDockerClient) ContainerList(context context.Context, options container.ListOptions) ([]types.Container, error) {
    args := mockDockerClient.Called(context, options)
    return args.Get(0).([]types.Container), args.Error(1)
}

func (mockDockerClient *MockDockerClient) ContainerStop(context context.Context, containerID string, options container.StopOptions) error {
    args := mockDockerClient.Called(context, containerID, options)
    return args.Error(0)
}

func (mockDockerClient *MockDockerClient) ContainerRemove(context context.Context, containerID string, options container.RemoveOptions) error {
    args := mockDockerClient.Called(context, containerID, options)
    return args.Error(0)
}

func (mockDockerClient *MockDockerClient) ImageList(context context.Context, options image.ListOptions) ([]image.Summary, error) {
    args := mockDockerClient.Called(context, options)
    return args.Get(0).([]image.Summary), args.Error(1)
}

func (mockDockerClient *MockDockerClient) ImageRemove(context context.Context, imageID string, options image.RemoveOptions) ([]image.DeleteResponse, error) {
    args := mockDockerClient.Called(context, imageID, options)
    return args.Get(0).([]image.DeleteResponse), args.Error(1)
}

func (mockDockerClient *MockDockerClient) NetworkList(context context.Context, options network.ListOptions) ([]network.Summary, error) {
    args := mockDockerClient.Called(context, options)
    return args.Get(0).([]network.Summary), args.Error(1)
}

func (mockDockerClient *MockDockerClient) NetworkRemove(context context.Context, networkID string) error {
    args := mockDockerClient.Called(context, networkID)
    return args.Error(0)
}

func (mockDockerClient *MockDockerClient) VolumeList(context context.Context, options volume.ListOptions) (volume.ListResponse, error) {
    args := mockDockerClient.Called(context, options)
    return args.Get(0).(volume.ListResponse), args.Error(1)
}
func (mockDockerClient *MockDockerClient) VolumeRemove(context context.Context, volumeID string, force bool) error {
    args := mockDockerClient.Called(context, volumeID, force)
    return args.Error(0)
}
