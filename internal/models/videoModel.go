package models

import "time"

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
