package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/njverse/go-fiber-hexagonal-template/config"
	"github.com/njverse/go-fiber-hexagonal-template/handler"
	"github.com/njverse/go-fiber-hexagonal-template/middleware"
)

func main() {
	config.InitEnv()
	config.InitTimeZone()

	db := config.InitDatabase()
	defer db.Close()

	fiberConfig := config.FiberConfig()
	// create a new Fiber instance
	app := fiber.New(fiberConfig)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Handlers.
	handler.FiberHandler(app, db) // Register Fiber's handler for app.

	// Start server
	log.Println("Server Running on Port: ", os.Getenv("SERVER_PORT"))
	app.Listen(fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")))
}
