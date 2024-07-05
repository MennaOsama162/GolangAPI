package routes

import (
	"library-management/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Author routes
	api.Post("/authors", controllers.CreateAuthor)
	api.Get("/authors", controllers.GetAuthors)
	api.Get("/authors/:id", controllers.GetAuthor)
	api.Put("/authors/:id", controllers.UpdateAuthor)
	api.Delete("/authors/:id", controllers.DeleteAuthor)

	// Book routes
	api.Post("/books", controllers.CreateBook)
	api.Get("/books", controllers.GetBooks)
	api.Get("/books/:id", controllers.GetBook)
	api.Put("/books/:id", controllers.UpdateBook)
	api.Delete("/books/:id", controllers.DeleteBook)
}
