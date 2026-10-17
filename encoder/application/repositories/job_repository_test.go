package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"testing"
	"encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"time"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryInsert(t *testing.T){
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()
	
	repo := repositories.NewVideoRepositoryDb(db)
	repo.Insert(video)

	job, err := domain.NewJob("path", "pending", video)
	require.Nil(t, err)
	jobRepo := repositories.NewJobRepositoryDb(db)
	jobRepo.Insert(job)

	j, err := jobRepo.FindById(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, j.ID)
	require.Equal(t, job.VideoID, j.VideoID)
}

func TestJobRepositoryUpdate(t *testing.T){
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()
	
	repo := repositories.NewVideoRepositoryDb(db)
	repo.Insert(video)

	job, err := domain.NewJob("path", "pending", video)
	require.Nil(t, err)
	jobRepo := repositories.NewJobRepositoryDb(db)
	jobRepo.Insert(job)

	job.Status = "completed"
	jobRepo.Update(job)

	j, err := jobRepo.FindById(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, j.ID)
	require.Equal(t, job.VideoID, j.VideoID)
	require.Equal(t, job.Status, j.Status)
}