package handler

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/guaychou/flusher/docs"
	"github.com/guaychou/flusher/internal/type"
)

// HealthHandler is function that for healthcheck this application
// Flusher Health Handler godoc
// @Summary Health Handler
// @Description Health handler for flusher
// @ID health-handler
// @Accept  json
// @Produce  json
// @Success 200 {object} types.HealthResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/health [get]
func HealthHanlder(c *fiber.Ctx) error {
	return c.Status(200).JSON(&types.HealthResponse{
		Success: true,
		Message: "OK",
		Version: os.Getenv("FLUSHER_VERSION"),
	})
}
