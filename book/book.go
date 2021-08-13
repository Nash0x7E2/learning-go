package book

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"learning_go/database"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)

	if book.Name == "" {
		c.Status(404).Send("No book found with the ID" + id)
		return
	}

	c.JSON(book)
}

func AddBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(500).Send("Unable to parse JSON data. Please review your request body")
		return
	}

	db.Create(&book)
	c.Status(200).JSON(&book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)

	if book.Name == "" {
		c.Status(404).Send("No book found with the ID " + id)
		return
	}
	db.Delete(&book)
	c.SendString("Book removed")
}
