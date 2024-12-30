package handlers

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// @Summary Get service health check
// @Description Responds with a message UP
// @Tags healthcheck
// @Produce json
// @Success 200 {string} string "ok"
// @Router /health [get]
func (h *HealthHandler) CheckHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "UP",
		"message": "Service is healthy",
	})
}
