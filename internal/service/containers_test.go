package service

import (
	"errors"
	"github.com/cloverchio/docker-nuke/internal/service/mocks"
	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	ContainerID           = "Container"
	ContainerList         = "ContainerList"
	ContainerStop         = "ContainerStop"
	ContainerRemove       = "ContainerRemove"
	ContainerErrorMessage = "Container Test Error"
)

func TestStopAllContainers(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)

	containerError := StopAllContainers(mockDockerClient)

	assert.NoError(t, containerError)
	mockDockerClient.AssertExpectations(t)
}

func TestStopAllContainers_ListError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{}, errors.New(ContainerErrorMessage))

	listError := StopAllContainers(mockDockerClient)

	assert.Error(t, listError)
	mockDockerClient.AssertNotCalled(t, ContainerStop, mock.Anything, mock.Anything, mock.Anything)
}

func TestStopAllContainers_StopError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(errors.New(ContainerErrorMessage))

	stopError := StopAllContainers(mockDockerClient)

	assert.Error(t, stopError)
	mockDockerClient.AssertExpectations(t)
}

func TestRemoveAllContainers(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ContainerRemove, mock.Anything, ContainerID, mock.Anything).Return(nil)

	containerError := RemoveAllContainers(mockDockerClient)

	assert.NoError(t, containerError)
	mockDockerClient.AssertExpectations(t)
}

func TestRemoveAllContainers_ListError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{}, errors.New(ContainerErrorMessage))

	listError := RemoveAllContainers(mockDockerClient)

	assert.Error(t, listError)
	mockDockerClient.AssertNotCalled(t, ContainerStop, mock.Anything, mock.Anything, mock.Anything)
}

func TestRemoveAllContainers_StopError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(errors.New(ContainerErrorMessage))

	stopError := RemoveAllContainers(mockDockerClient)

	assert.Error(t, stopError)
	mockDockerClient.AssertNotCalled(t, ContainerRemove, mock.Anything, ContainerID, mock.Anything)
}

func TestRemoveAllContainers_RemoveError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ContainerRemove, mock.Anything, ContainerID, mock.Anything).Return(errors.New(ContainerErrorMessage))

	removeError := RemoveAllContainers(mockDockerClient)

	assert.Error(t, removeError)
	mockDockerClient.AssertExpectations(t)
}
