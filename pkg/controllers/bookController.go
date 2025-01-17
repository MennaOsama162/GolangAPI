package controllers

import (
	"time"

	"library-management/pkg/config"
	"library-management/pkg/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateBookInput struct {
	Title         string `json:"title" validate:"required"`
	ISBN          string `json:"isbn" validate:"required"`
	PublishedDate string `json:"published_date" validate:"required"`
	AuthorID      uint   `json:"author_id" validate:"required"`
}

/////////////////////////////////////////////////Create ////////////////////////////////////////////////////////////

func CreateBook(c *fiber.Ctx) error {
	var input CreateBookInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	publishedDate, err := time.Parse("2006-01-02", input.PublishedDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}

	// Check for duplicate ISBN
	var existingBook models.Book
	if err := config.DB.Where("isbn = ?", input.ISBN).First(&existingBook).Error; err == nil {
		// Book with the same ISBN already exists
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Book with this ISBN already exists"})
	}

	book := models.Book{Title: input.Title, ISBN: input.ISBN, PublishedDate: publishedDate, AuthorID: input.AuthorID}
	if err := config.DB.Create(&book).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(book)
}

///////////////////////////////////////////////get All ///////////////////////////////////////////////////

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.DB.Preload("Author").Find(&books)
	return c.JSON(books)
}

///////////////////////////////////////////////get one ///////////////////////////////////////////////////

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.Preload("Author").First(&book, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

///////////////////////////////////////////////update ///////////////////////////////////////////////////

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	var input CreateBookInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	publishedDate, err := time.Parse("2006-01-02", input.PublishedDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}

	book.Title = input.Title
	book.ISBN = input.ISBN
	book.PublishedDate = publishedDate
	book.AuthorID = input.AuthorID
	config.DB.Save(&book)
	return c.JSON(book)
}

///////////////////////////////////////////////delete ///////////////////////////////////////////////////

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	config.DB.Delete(&book)
	return c.SendStatus(http.StatusNoContent)
}

//////////////////////////////////////////////Soft delete //////////////////////////////////////////////

func SoftDeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Set DeletedAt to the current timestamp
	book.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if result := config.DB.Save(&book); result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to soft delete book"})
	}

	return c.SendStatus(http.StatusNoContent)
}

/////////////////////////////////////////Search By Title///////////////////////////////////

func SearchBookByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	var books []models.Book
	config.DB.Preload("Author").Where("title LIKE ?", "%"+title+"%").Find(&books)
	return c.JSON(books)
}
