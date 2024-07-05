package routes

import (
	"library-management/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	authorRoutes := api.Group("/authors")
	authorRoutes.Post("/", controllers.CreateAuthor)
	authorRoutes.Get("/", controllers.GetAuthors)
	authorRoutes.Get("/:id", controllers.GetAuthor)
	authorRoutes.Put("/:id", controllers.UpdateAuthor)
	authorRoutes.Delete("/:id", controllers.DeleteAuthor)

	bookRoutes := api.Group("/books")
	bookRoutes.Post("/", controllers.CreateBook)
	bookRoutes.Get("/", controllers.GetBooks)
	bookRoutes.Get("/:id", controllers.GetBook)
	bookRoutes.Put("/:id", controllers.UpdateBook)
	bookRoutes.Delete("/:id", controllers.DeleteBook)
}
