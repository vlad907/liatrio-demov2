package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // spin up a new Fiber instance (lightweight web framework)

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.SetContentType("application/json; charset=utf-8")                            // tell clients we’re returning JSON
		clean := fmt.Sprintf(`{"message":"My name is John Doe","timestamp":%d}`, time.Now().UnixMilli()) // craft the payload with a Unix ms timestamp
		return c.SendString(clean)
	})

	port := os.Getenv("PORT") // Cloud Run sets PORT; fallback keeps things working locally
	if port == "" {
		port = "80" // default to 80 so the container and tests stay aligned
	}

	log.Printf("Server starting on port %s...", port) // helpful startup log for container logs
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err) // crash fast if the listener can’t bind
	}
}
