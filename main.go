package main

import (
	"context"
	"log"
	"social/app"
	"social/opentelemetry"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := fiber.New()
	tracerProvider := opentelemetry.InitTracer()
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	server := app.NewServer(fiberApp)

	server.Start()
}
