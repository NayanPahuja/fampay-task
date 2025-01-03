package models

import "time"

// Video struct represents a video object
// @Description Video object representing a video
type Video struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"type:varchar(255);unique;not null"`
	VideoID      string    `gorm:"unique"` // YouTube video ID
	Description  string    `gorm:"type:text"`
	PublishedAt  time.Time `gorm:"index"`
	ThumbnailURL string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// PaginationResponse struct represents a video object
// @Description Video object representing a video
type PaginationResponse struct {
	Videos     []Video `json:"videos"`      // The list of videos
	NextCursor string  `json:"next_cursor"` // Cursor for the next page
}
