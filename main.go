package main

import (
    "bytes"
    "encoding/json"           // <— add this
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
	"strconv"
)

func main() {
    app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal, // ensures minified JSON
    })

    app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		c.Set("Cache-Control", "no-cache")
		
		response := `{"message":"My name is John Doe","timestamp":` + 
			strconv.FormatInt(time.Now().UnixMilli(), 10) + `}`
		
		return c.SendString(response)
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