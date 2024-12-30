// workflows/youtube.go
package workflows

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/NayanPahuja/fam-bcknd-test/db"
	"github.com/NayanPahuja/fam-bcknd-test/internal/models"
	"github.com/NayanPahuja/fam-bcknd-test/internal/utils"
)

const YouTubeAPIURL = "https://www.googleapis.com/youtube/v3/search"

var keyManager *utils.APIKeyManager

// init function to initialize the key manager
func init() {
	var err error
	keyManager, err = utils.NewAPIKeyManager()
	if err != nil {
		log.Fatalf("Failed to initialize API key manager: %v", err)
	}
}
func YouTubeActivity(ctx context.Context, searchQuery string) error {
	currentTime := time.Now()
	dayBeforeYesterday := currentTime.AddDate(0, 0, -2)
	publishedAfter := dayBeforeYesterday.Format(time.RFC3339)

	baseURL, err := url.Parse(YouTubeAPIURL)
	if err != nil {
		return fmt.Errorf("error parsing YouTube API URL: %w", err)
	}

	keyManager.Reset()
	videoCache := make(map[string]bool)
	pageCount := 0

	for {
		currentKey := keyManager.GetCurrentKey()
		params := url.Values{}
		params.Add("part", "snippet")
		params.Add("type", "video")
		params.Add("order", "date")
		params.Add("q", searchQuery)
		params.Add("key", currentKey)
		params.Add("maxResults", "50")
		params.Add("publishedAfter", publishedAfter)

		log.Printf("Using YouTube API key: %v", currentKey)

		// Retry logic for handling quota exceeded errors
		var ytResponse utils.YouTubeResponse
		success := false

		for !success {
			baseURL.RawQuery = params.Encode()
			currentURL := baseURL.String()

			req, err := http.NewRequestWithContext(ctx, "GET", currentURL, nil)
			if err != nil {
				return fmt.Errorf("error creating request: %w", err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("error making YouTube API request: %w", err)
			}

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode == 403 && utils.IsQuotaExceededError(resp.Body) {
					resp.Body.Close()

					// Try getting next API key
					if nextKey, ok := keyManager.NextKey(); ok {
						log.Printf("Quota exceeded. Switching to next API key.")
						params.Set("key", nextKey)
						log.Printf("Using YouTube API key: %v", nextKey)
						continue
					} else {
						return fmt.Errorf("all API keys have exceeded their quota")
					}
				}

				body := make([]byte, 1024)
				resp.Body.Read(body)
				resp.Body.Close()
				return fmt.Errorf("failed to fetch data from YouTube API: status %d, body: %s", resp.StatusCode, string(body))
			}

			if err := json.NewDecoder(resp.Body).Decode(&ytResponse); err != nil {
				resp.Body.Close()
				return fmt.Errorf("error decoding YouTube API response: %w", err)
			}
			resp.Body.Close()
			success = true
		}

		// Process videos
		for _, item := range ytResponse.Items {
			if videoCache[item.ID.VideoID] {
				log.Println("Skipping as this video already exists in cache!")
				continue
			}

			publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				log.Printf("Error parsing published date: %v", err)
				continue
			}

			video := models.Video{
				VideoID:      item.ID.VideoID,
				Title:        item.Snippet.Title,
				Description:  item.Snippet.Description,
				PublishedAt:  publishedAt,
				ThumbnailURL: item.Snippet.Thumbnails.Default.URL,
			}

			if err := db.DB.Create(&video).Error; err != nil {
				log.Printf("Error saving video to database: %v", err)
				continue
			}

			videoCache[item.ID.VideoID] = true
		}

		log.Printf("Successfully fetched and saved %d videos", len(ytResponse.Items))

		pageCount++
		if ytResponse.NextPageToken != "" && pageCount < 10 {
			params.Set("pageToken", ytResponse.NextPageToken)
		} else {
			break
		}
	}

	return nil
}
