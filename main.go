package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
    c.Response().Header.SetContentType("application/json; charset=utf-8")
    clean := fmt.Sprintf(`{"message":"My name is John Doe","timestamp":%d}`, time.Now().UnixMilli())
    return c.SendString(clean)
	})

    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }

    log.Printf("Server starting on port %s...", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}