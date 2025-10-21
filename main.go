package main

import (
    "encoding/json"           // <— add this
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal, // ensures minified JSON
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message":   "My name is John Doe",
            "timestamp": time.Now().UnixMilli(),
        })
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