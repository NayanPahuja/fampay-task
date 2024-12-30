package repositories

import (
	"log"

	"github.com/NayanPahuja/fam-bcknd-test/internal/models"
	"gorm.io/gorm"
)

type VideoRepository interface {
	GetVideosByPagination(limit int, offset int) ([]models.Video, error)
}

type videoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) VideoRepository {
	return &videoRepository{
		db: db,
	}
}

func (repo *videoRepository) GetVideosByPagination(limit int, offset int) ([]models.Video, error) {
	var videos []models.Video
	err := repo.db.Order("published_at DESC").Limit(limit).Offset(offset).Find(&videos).Error

	if err != nil {
		log.Println("Unable to get videos from the database!")
		return nil, err
	}

	return videos, err
}
