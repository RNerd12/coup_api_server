package api

import (
	"github.com/gofiber/fiber/v2"
)

func HealthPing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "OK",
	})
}
