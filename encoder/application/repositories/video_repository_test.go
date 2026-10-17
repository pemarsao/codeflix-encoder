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

func TestVideoRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()
	
	repo := repositories.NewVideoRepositoryDb(db)
	repo.Insert(video)

	v, err := repo.FindById(video.ID)
	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, video.ID, v.ID)
}