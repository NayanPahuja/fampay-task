package routes

import (
	"github.com/NayanPahuja/fam-bcknd-test/internal/handlers"
	"github.com/NayanPahuja/fam-bcknd-test/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	videoService := services.NewVideoService()
	videoHandler := handlers.NewVideoHandler(videoService)
	api := app.Group("/api/v1")
	api.Get("/videos", videoHandler.GetPaginatedVideos)
	api.Get("/videosv2", videoHandler.GetPaginatedVideosUsingCursor)
	healthHandler := handlers.NewHealthHandler()
	app.Get("/health", healthHandler.CheckHealth)
}
