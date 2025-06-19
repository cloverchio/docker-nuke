package service

import (
	"errors"
	"github.com/cloverchio/docker-nuke/internal/service/mocks"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	VolumeID           = "Volume"
	VolumeList         = "VolumeList"
	VolumeRemove       = "VolumeRemove"
	VolumeErrorMessage = "Volume Test Error"
)

func TestRemoveAllVolumes(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	// mock container stop check
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	// volume checks
	mockDockerClient.On(VolumeList, mock.Anything, mock.Anything).Return(volume.ListResponse{
		Volumes: []*volume.Volume{{Name: VolumeID}},
	}, nil)
	mockDockerClient.On(VolumeRemove, mock.Anything, VolumeID, true).Return(nil)

	volumeError := RemoveAllVolumes(mockDockerClient)

	assert.NoError(t, volumeError)
	mockDockerClient.AssertExpectations(t)
}

func TestRemoveAllVolumes_ContainerStopError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(errors.New(VolumeErrorMessage))

	stopError := RemoveAllVolumes(mockDockerClient)

	assert.Error(t, stopError)
	mockDockerClient.AssertNotCalled(t, VolumeList, mock.Anything, mock.Anything)
	mockDockerClient.AssertNotCalled(t, VolumeRemove, mock.Anything, VolumeID, true)
}

func TestRemoveAllVolumes_ListError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(VolumeList, mock.Anything, mock.Anything).Return(volume.ListResponse{
		Volumes: []*volume.Volume{{Name: VolumeID}},
	}, errors.New(VolumeErrorMessage))

	listError := RemoveAllVolumes(mockDockerClient)

	assert.Error(t, listError)
	mockDockerClient.AssertNotCalled(t, VolumeRemove, mock.Anything, VolumeID, true)
}

func TestRemoveAllVolumes_RemoveError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(VolumeList, mock.Anything, mock.Anything).Return(volume.ListResponse{
		Volumes: []*volume.Volume{{Name: VolumeID}},
	}, nil)
	mockDockerClient.On(VolumeRemove, mock.Anything, VolumeID, true).Return(errors.New(VolumeErrorMessage))

	removeError := RemoveAllVolumes(mockDockerClient)

	assert.Error(t, removeError)
	mockDockerClient.AssertExpectations(t)
}
