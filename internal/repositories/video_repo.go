package repositories

import (
	"fmt"
	"log"

	"github.com/NayanPahuja/fam-bcknd-test/internal/models"
	"github.com/NayanPahuja/fam-bcknd-test/internal/utils"
	"gorm.io/gorm"
)

type VideoRepository interface {
	GetVideosByPagination(limit int, offset int) ([]models.Video, error)
	GetVideosByCursor(encodedCursor string, limit int) ([]models.Video, string, error)
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

func (repo *videoRepository) GetVideosByCursor(encodedCursor string, limit int) ([]models.Video, string, error) {
	var videos []models.Video
	var cursor utils.Cursor
	var err error

	// Decode cursor if provided
	if encodedCursor != "" {
		cursor, err = utils.DecodeCursor(encodedCursor)
		if err != nil {
			log.Printf("Failed to decode cursor: %v", err)
			return nil, "", fmt.Errorf("invalid cursor format")
		}
	}
	log.Printf("Decoded cursor: %+v", cursor)
	// Build query
	query := repo.db.Order("published_at DESC").Limit(limit)

	// Apply cursor filter
	if cursor.PublishedAt != "" {
		query = query.Where("published_at < ?", cursor.PublishedAt)
	}
	log.Printf("Query: %+v, Cursor: %+v", query, cursor)
	// Execute query
	err = query.Find(&videos).Error
	if err != nil {
		log.Printf("Database query failed: %v", err)
		return nil, "", err
	}

	// Generate next cursor if there are more results
	var nextCursor string
	if len(videos) == limit {
		lastVideo := videos[len(videos)-1]
		nextCursor = utils.EncodeCursor(lastVideo.PublishedAt.String())
	}
	log.Printf("Next cursor: %s", nextCursor)
	return videos, nextCursor, nil
}
