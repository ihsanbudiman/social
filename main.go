package main

import (
	"social/app"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := fiber.New()

	server := app.NewServer(fiberApp)

	server.Start()
}
