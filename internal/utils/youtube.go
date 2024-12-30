// utils/youtube.go
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/NayanPahuja/fam-bcknd-test/config"
)

type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			PublishedAt string `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

type YouTubeErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

type APIKeyManager struct {
	keys     []string
	current  int
	mu       sync.Mutex
	attempts int
}

// NewAPIKeyManager creates a new API key manager using keys from config
func NewAPIKeyManager() (*APIKeyManager, error) {
	keys := config.Envs.YouTubeAPIKeys
	if len(keys) == 0 {
		return nil, fmt.Errorf("no YouTube API keys configured")
	}

	return &APIKeyManager{
		keys:     keys,
		current:  0,
		attempts: 0,
	}, nil
}

func (m *APIKeyManager) GetCurrentKey() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.keys[m.current]
}

func (m *APIKeyManager) NextKey() (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.attempts++
	if m.attempts >= len(m.keys) {
		// Reset attempts counter and return false to indicate we've tried all keys
		m.attempts = 0
		return "", false
	}

	m.current = (m.current + 1) % len(m.keys)
	return m.keys[m.current], true
}

func (m *APIKeyManager) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.current = 0
	m.attempts = 0
}

func IsQuotaExceededError(responseBody io.Reader) bool {
	var errResp YouTubeErrorResponse
	body, err := io.ReadAll(responseBody)
	if err != nil {
		return false
	}

	if err := json.Unmarshal(body, &errResp); err != nil {
		return false
	}

	// Check if it's a quota exceeded error
	if errResp.Error.Code == 403 {
		for _, err := range errResp.Error.Errors {
			if err.Domain == "youtube.quota" && err.Reason == "quotaExceeded" {
				return true
			}
		}
	}

	return false
}
