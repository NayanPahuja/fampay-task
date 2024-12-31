package services

import (
	"github.com/NayanPahuja/fam-bcknd-test/db"
	"github.com/NayanPahuja/fam-bcknd-test/internal/models"
	"github.com/NayanPahuja/fam-bcknd-test/internal/repositories"
)

type VideoService interface {
	GetPaginatedVideos(limit, offset int) ([]models.Video, error)
	GetPaginatedVideosUsingCursor(encodedCursor string, limit int) ([]models.Video, string, error)
}

type videoService struct {
	repo repositories.VideoRepository
}

func NewVideoService() VideoService {
	return &videoService{repo: repositories.NewVideoRepository(db.DB)}
}

func (s *videoService) GetPaginatedVideos(limit int, offset int) ([]models.Video, error) {
	return s.repo.GetVideosByPagination(limit, offset)
}

func (s *videoService) GetPaginatedVideosUsingCursor(encodedCursor string, limit int) ([]models.Video, string, error) {
	return s.repo.GetVideosByCursor(encodedCursor, limit)
}
