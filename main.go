package main

import (
	"log"

	"library-management/pkg/config"
	"library-management/pkg/models"
	"library-management/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Connect to MySQL
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Author{}, &models.Book{})

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
