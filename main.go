package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"learning_go/book"
	"learning_go/database"
	"log"
)

func configureRoutes(app *fiber.App) {
	app.Get("api/v1/books", book.GetBooks)
	app.Get("api/v1/books/:id", book.GetBook)
	app.Post("api/v1/books", book.AddBook)
	app.Delete("api/v1/books/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	conn := database.InitDatabase()
	conn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated using &book.Book 📕")
	configureRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
