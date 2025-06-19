package service

import (
	"errors"
	"github.com/cloverchio/docker-nuke/internal/service/mocks"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	ImageID           = "Image"
	DanglingImageID   = "Dangling Image"
	RepoTag           = "Repo Tag"
	ImageList         = "ImageList"
	ImageRemove       = "ImageRemove"
	ImageErrorMessage = "Image Test Error"
)

func TestRemoveDanglingImages(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	// mock container stop check
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	// image checks
	mockDockerClient.On(ImageList, mock.Anything, mock.Anything).Return([]image.Summary{
		{ID: ImageID, RepoTags: []string{RepoTag}},
		// dangling images have no repo tags...
		{ID: DanglingImageID},
	}, nil)
	mockDockerClient.On(ImageRemove, mock.Anything, DanglingImageID, mock.Anything).Return([]image.DeleteResponse{}, nil)

	imageError := RemoveDanglingImages(mockDockerClient)

	assert.NoError(t, imageError)
	// check that the non dangling image wasn't deleted
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, ImageID, mock.Anything)
}

func TestRemoveDanglingImages_ContainerStopError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(errors.New(ImageErrorMessage))

	stopError := RemoveDanglingImages(mockDockerClient)

	assert.Error(t, stopError)
	mockDockerClient.AssertNotCalled(t, ImageList, mock.Anything, mock.Anything)
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, mock.Anything, mock.Anything)
}

func TestRemoveDanglingImages_ListError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ImageList, mock.Anything, mock.Anything).Return([]image.Summary{
		{ID: ImageID, RepoTags: []string{RepoTag}},
		{ID: DanglingImageID},
	}, errors.New(ImageErrorMessage))

	listError := RemoveDanglingImages(mockDockerClient)

	assert.Error(t, listError)
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, mock.Anything, mock.Anything)
}

func TestRemoveDanglingImages_RemoveError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ImageList, mock.Anything, mock.Anything).Return([]image.Summary{
		{ID: ImageID, RepoTags: []string{RepoTag}},
		{ID: DanglingImageID},
	}, nil)
	mockDockerClient.On(ImageRemove, mock.Anything, DanglingImageID, mock.Anything).Return([]image.DeleteResponse{}, errors.New(ImageErrorMessage))

	removeError := RemoveDanglingImages(mockDockerClient)

	assert.Error(t, removeError)
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, ImageID, mock.Anything)
}

func TestRemoveAllImages(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	// mock container stop check
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ImageList, mock.Anything, mock.Anything).Return([]image.Summary{
		{ID: ImageID, RepoTags: []string{RepoTag}},
		{ID: DanglingImageID},
	}, nil)
	mockDockerClient.On(ImageRemove, mock.Anything, DanglingImageID, mock.Anything).Return([]image.DeleteResponse{}, nil)
	mockDockerClient.On(ImageRemove, mock.Anything, ImageID, mock.Anything).Return([]image.DeleteResponse{}, nil)

	imageError := RemoveAllImages(mockDockerClient)

	assert.NoError(t, imageError)
	mockDockerClient.AssertExpectations(t)
}

func TestRemoveAllImages_ContainerStopError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(errors.New(ImageErrorMessage))

	stopError := RemoveAllImages(mockDockerClient)

	assert.Error(t, stopError)
	mockDockerClient.AssertNotCalled(t, ImageList, mock.Anything, mock.Anything)
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, mock.Anything, mock.Anything)
}

func TestRemoveAllImages_ListError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ImageList, mock.Anything, mock.Anything).Return([]image.Summary{
		{ID: ImageID, RepoTags: []string{RepoTag}},
		{ID: DanglingImageID},
	}, errors.New(ImageErrorMessage))

	listError := RemoveAllImages(mockDockerClient)

	assert.Error(t, listError)
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, mock.Anything, mock.Anything)
}

func TestRemoveAllImages_RemoveError(t *testing.T) {
	mockDockerClient := new(mocks.MockDockerClient)
	mockDockerClient.On(ContainerList, mock.Anything, mock.Anything).Return([]types.Container{
		{ID: ContainerID},
	}, nil)
	mockDockerClient.On(ContainerStop, mock.Anything, ContainerID, mock.Anything).Return(nil)
	mockDockerClient.On(ImageList, mock.Anything, mock.Anything).Return([]image.Summary{
		{ID: ImageID, RepoTags: []string{RepoTag}},
		{ID: DanglingImageID},
	}, nil)
	mockDockerClient.On(ImageRemove, mock.Anything, ImageID, mock.Anything).Return([]image.DeleteResponse{}, errors.New(ImageErrorMessage))

	removeError := RemoveAllImages(mockDockerClient)

	assert.Error(t, removeError)
	mockDockerClient.AssertNotCalled(t, ImageRemove, mock.Anything, DanglingImageID, mock.Anything)
}
