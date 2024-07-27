package main

import (
	"github.com/cnc-csku/go-clean-arch-example/internal/adapters/memory"
	"github.com/cnc-csku/go-clean-arch-example/internal/adapters/rest"
	usecases "github.com/cnc-csku/go-clean-arch-example/usecases/book"
	"github.com/gofiber/fiber/v2"
)

func main() {
	bookRepo := memory.NewBookMemory()
	bookService := usecases.NewBookService(bookRepo)
	bookHandler := rest.NewBookRestHandler(bookService)

	app := fiber.New()

	app.Get("/books", bookHandler.GetAllBooks)
	app.Post("/books", bookHandler.CreateBook)

	app.Listen(":3000")
}