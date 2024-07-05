package controllers

import (
	"library-management/pkg/config"
	"library-management/pkg/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type CreateAuthorInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func CreateAuthor(c *fiber.Ctx) error {
	var input CreateAuthorInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	author := models.Author{Name: input.Name, Email: input.Email}
	config.DB.Create(&author)
	return c.Status(http.StatusCreated).JSON(author)
}

func GetAuthors(c *fiber.Ctx) error {
	var authors []models.Author
	config.DB.Find(&authors)
	return c.JSON(authors)
}

func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	if result := config.DB.First(&author, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
	}
	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	if result := config.DB.First(&author, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
	}

	var input CreateAuthorInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	author.Name = input.Name
	author.Email = input.Email
	config.DB.Save(&author)
	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	if result := config.DB.First(&author, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Author not found"})
	}

	config.DB.Delete(&author)
	return c.SendStatus(http.StatusNoContent)
}
