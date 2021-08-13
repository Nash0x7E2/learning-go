package main

import (
	"github.com/gofiber/fiber"
	"learning_go/book"
)

func configureRoutes(app *fiber.App) {
	app.Get("api/v1/books", book.GetBooks)
	app.Get("api/v1/books/:id", book.GetBook)
	app.Post("api/v1/books", book.AddBook)
	app.Delete("api/v1/books/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	configureRoutes(app)

	app.Listen(":3000")
}
