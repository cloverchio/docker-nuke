package service

import (
	"errors"
	"github.com/cloverchio/docker-nuke/internal/service/mocks"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	NetworkID           = "Network"
	NetworkList         = "NetworkList"
	NetworkRemove       = "NetworkRemove"
	NetworkErrorMessage = "Network Test Error"
)

func TestRemoveAllNetworks(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	// mock container stop check
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	// network checks
	mockDockerClient.On(NetworkList, mock.Anything, mock.Anything).Return([]network.Summary{
		{ID: NetworkID},
	}, nil)
	mockDockerClient.On(NetworkRemove, mock.Anything, NetworkID).Return(nil)

	networkError := RemoveAllNetworks(mockDockerClient)

	assert.NoError(t, networkError)
	mockDockerClient.AssertExpectations(t)
}

func TestRemoveAllNetworks_ContainerStopError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(errors.New(NetworkErrorMessage))

	stopError := RemoveAllNetworks(mockDockerClient)

	assert.Error(t, stopError)
	mockDockerClient.AssertNotCalled(t, NetworkList, mock.Anything, mock.Anything)
	mockDockerClient.AssertNotCalled(t, NetworkRemove, mock.Anything, NetworkID)
}

func TestRemoveAllNetworks_ListError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(NetworkList, mock.Anything, mock.Anything).Return([]network.Summary{
		{ID: NetworkID},
	}, errors.New(NetworkErrorMessage))

	listError := RemoveAllNetworks(mockDockerClient)

	assert.Error(t, listError)
	mockDockerClient.AssertNotCalled(t, NetworkRemove, mock.Anything, NetworkID)
}

func TestRemoveAllNetworks_RemoveError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(NetworkList, mock.Anything, mock.Anything).Return([]network.Summary{
		{ID: NetworkID},
	}, nil)
	mockDockerClient.On(NetworkRemove, mock.Anything, NetworkID).Return(errors.New(NetworkErrorMessage))

	removeError := RemoveAllNetworks(mockDockerClient)

	assert.Error(t, removeError)
	mockDockerClient.AssertExpectations(t)
}
