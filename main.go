package main

import (
	"streaming/config/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading env")
	}

	database.ConnectDB()

	app := fiber.New()

	subscribers := app.Group("/subscribers")

	subscribers.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SAAS")
	})

	app.Listen(":8080")
}
