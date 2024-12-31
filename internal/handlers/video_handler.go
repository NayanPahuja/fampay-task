package handlers

import (
	"strconv"

	_ "github.com/NayanPahuja/fam-bcknd-test/internal/models"
	"github.com/NayanPahuja/fam-bcknd-test/internal/services"
	"github.com/gofiber/fiber/v2"
)

type VideoHandler struct {
	service services.VideoService
}

func NewVideoHandler(service services.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}

// GetVideos godoc
// @Summary Get list of videos
// @Description Get a list of videos with pagination support
// @Tags videos
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit per page"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array}  models.Video "Ok"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/videos [get]
func (h *VideoHandler) GetPaginatedVideos(c *fiber.Ctx) error {
	//Parse query parameters

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil || offset < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid offset parameter",
		})
	}

	// Fetch paginated videos
	videos, err := h.service.GetPaginatedVideos(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch videos",
		})
	}

	return c.JSON(videos)
}
