package main

import (
	"log"

	"coup_api_server/internal/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/health-ping", api.HealthPing)
	log.Fatal(app.Listen(":3000"))
}
