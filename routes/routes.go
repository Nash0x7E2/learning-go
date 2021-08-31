package routes

import (
	"github.com/gofiber/fiber"
	"learning_go/book"
	"learning_go/database"
)

func GetBooks(c *fiber.Ctx) {
	db := database.DB.Db
	var books []book.Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB.Db
	var book book.Book
	db.Find(&book, id)

	if book.Name == "" {
		c.Status(404).Send("No book found with the ID" + id)
		return
	}

	c.JSON(book)
}

func AddBook(c *fiber.Ctx) {
	db := database.DB.Db
	book := new(book.Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(500).Send("Unable to parse JSON data. Please review your request body")
		return
	}

	db.Create(&book)
	c.Status(200).JSON(&book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB.Db
	var book book.Book
	db.First(&book, id)

	if book.Name == "" {
		c.Status(404).Send("No book found with the ID " + id)
		return
	}
	db.Delete(&book)
	c.SendString("Book removed")
}
