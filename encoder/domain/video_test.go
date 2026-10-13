package domain_test

import (
	"encoder/domain"
	"testing"
	"github.com/stretchr/testify/require"
	"time"
	uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "invalid-uuid"
	video.ResourceID = "valid-resource-id"
	video.FilePath = "valid-file-path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestValidateResourceIDIsNull(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "invalid-uuid"
	video.FilePath = "valid-file-path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestValidateFilePathIsNull(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "invalid-uuid"
	video.ResourceID = "valid-resource-id"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidate(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "valid-resource-id"
	video.FilePath = "valid-file-path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}