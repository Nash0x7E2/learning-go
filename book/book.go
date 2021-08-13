package book

import "github.com/gofiber/fiber"

type Book struct {
	name   string `json:"name"`
	author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func GetBooks(c *fiber.Ctx) {
	c.Send("Hello books")
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	c.Send("Hello " + id)
}

func AddBook(c *fiber.Ctx) {
	c.Send("Hello add books")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Bye books")
}
