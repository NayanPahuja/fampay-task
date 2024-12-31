package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Cursor struct {
	PublishedAt string `json:"published_at"`
}

func EncodeCursor(publishedAt string) string {
	// Parse the publishedAt using the custom format that matches the input
	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 IST", publishedAt)
	if err != nil {
		log.Printf("Failed to parse publishedAt for encoding: %v", err)
		return ""
	}

	// Format it to RFC3339
	cursor := Cursor{PublishedAt: parsedTime.Format(time.RFC3339)}
	cursorBytes, err := json.Marshal(cursor)
	if err != nil {
		log.Printf("Failed to encode cursor: %v", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(cursorBytes)
}

func DecodeCursor(encodedCursor string) (Cursor, error) {
	var cursor Cursor
	cursorBytes, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return cursor, fmt.Errorf("invalid cursor format: %v", err)
	}
	err = json.Unmarshal(cursorBytes, &cursor)
	if err != nil {
		return cursor, fmt.Errorf("invalid cursor data: %v", err)
	}

	// Validate the PublishedAt timestamp
	_, err = time.Parse(time.RFC3339, cursor.PublishedAt)
	if err != nil {
		return cursor, fmt.Errorf("invalid timestamp format in cursor: %v", err)
	}

	return cursor, nil
}
