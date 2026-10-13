package domain_test

import (
	"encoder/domain"
	"testing"
	"github.com/stretchr/testify/require"
	"time"
	uuid "github.com/satori/go.uuid"
)

func TestValidateNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "valid-resource-id"
	video.FilePath = "valid-file-path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("valid-output-bucket-path", "valid-status", video)

	require.Nil(t, err)
	require.NotNil(t, job)
}