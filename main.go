package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"learning_go/database"
	"learning_go/routes"
	"log"
)

func configureRoutes(app *fiber.App) {
	app.Get("api/v1/books", routes.GetBooks)
	app.Get("api/v1/books/:id", routes.GetBook)
	app.Post("api/v1/books", routes.AddBook)
	app.Delete("api/v1/books/:id", routes.DeleteBook)
}

func main() {
	database.InitDatabase()

	app := fiber.New()
	app.Use(recover.New())

	configureRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
