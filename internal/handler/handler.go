package handler

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

// HealthHandler is function that for healthcheck this application
func HealthHanlder(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "OK",
		"version": os.Getenv("FLUSHER_VERSION"),
	})
}
