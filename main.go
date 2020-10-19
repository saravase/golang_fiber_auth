package main

import (
	"golang_fiber_auth/auth-api/database"
	"golang_fiber_auth/auth-api/handler"
	"golang_fiber_auth/auth-api/router"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {

	// New creates a new plant-api Logger.
	logger := log.New(os.Stdout, "auth-api", log.LstdFlags)

	// Initialize server configuration
	options := fiber.Config{
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	// Load .env data
	godotenv.Load()

	// Initialize database
	db := database.InitDatabase()
	
	// Initialize auth properties
	authHandler := handler.NewAuth(logger, db)

	// Initialize fiber app
	app := fiber.New(options)
	
	// Initialize router
	router.InitRoutes(app, authHandler)

	// Initialize listen port
	app.Listen(":9090")

}
