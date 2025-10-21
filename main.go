package main

import (
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message":   "My name is Mike Hunt",
            "timestamp": time.Now().Unix(),
        })
    })

    // Log and handle errors if the app fails to start
    if err := app.Listen(":80"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}