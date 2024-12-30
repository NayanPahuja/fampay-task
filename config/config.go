package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	TemporalHost         string
	TaskQueueName        string
	PublicHost           string
	Port                 string
	DBUser               string
	DBPassword           string
	DBHost               string
	DBPort               string
	DBName               string
	YouTubeAPIKeys       []string
	FetchIntervalSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		TemporalHost:         getEnv("TEMPORAL_HOST", "localhost:7233"),
		TaskQueueName:        getEnv("TASK_QUEUE_NAME", "YouTubeFetcherTaskQueue"),
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBUser:               getEnv("DB_USER", "postgres"),
		DBPassword:           getEnv("DB_PASSWORD", "password"),
		DBHost:               getEnv("DB_HOST", "127.0.0.1"),
		DBPort:               getEnv("DB_PORT", "5432"),
		DBName:               getEnv("DB_NAME", "youtube"),
		YouTubeAPIKeys:       getEnvAsSlice("YOUTUBE_API_KEYS", []string{"dummy-api-key"}),
		FetchIntervalSeconds: getEnvAsInt("FETCH_INTERVAL_SECONDS", 30),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return parsed
		}
	}
	return fallback
}

func getEnvAsSlice(key string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		return splitCommaSeparated(value)
	}
	return fallback
}

func splitCommaSeparated(value string) []string {
	return removeEmptyStrings(split(value, ','))
}

func split(value string, sep rune) []string {
	return strings.FieldsFunc(value, func(r rune) bool { return r == sep })
}

func removeEmptyStrings(slice []string) []string {
	var result []string
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
