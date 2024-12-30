package routes

import (
	"github.com/NayanPahuja/fam-bcknd-test/internal/handlers"
	"github.com/NayanPahuja/fam-bcknd-test/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	videoService := services.NewVideoService()
	videoHandler := handlers.NewVideoHandler(videoService)

	api := app.Group("/api")
	api.Get("/videos", videoHandler.GetPaginatedVideos)

	healthHandler := handlers.NewHealthHandler()
	app.Get("/health", healthHandler.CheckHealth)
}
