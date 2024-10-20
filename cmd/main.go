package main

import (
	"book-list/data"
	"book-list/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	repo := data.NewInMemoryRepository()

	// Create a test user
	repo.Users[1] = &data.User{ID: 1, Name: "John Doe"}

	bookHandler := handlers.NewBookHandler(repo)

	app := fiber.New()

	app.Get("/", bookHandler.GetBooks)
	app.Post("/books", bookHandler.AddBook)
	app.Delete("/books/:id", bookHandler.DeleteBook)

	log.Fatal(app.Listen(":3000"))
}
