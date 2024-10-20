package handlers

import (
	"book-list/data"
	"book-list/templates"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type BookHandler struct {
	repo data.Repository
}

func NewBookHandler(repo data.Repository) *BookHandler {
	return &BookHandler{repo: repo}
}

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	user, err := h.repo.GetUser(1) // Hardcoded user ID for simplicity
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	c.Set("content-type", fiber.MIMETextHTMLCharsetUTF8)
	return templates.BookList(user).Render(c.Context(), c.Response().BodyWriter())
}

func (h *BookHandler) AddBook(c *fiber.Ctx) error {
	title := c.FormValue("title")
	author := c.FormValue("author")

	book := data.Book{Title: title, Author: author}
	err := h.repo.AddBook(1, book) // Hardcoded user ID for simplicity
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return h.GetBooks(c) // Re-render the entire book list
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
	}

	err = h.repo.DeleteBook(1, bookID) // Hardcoded user ID for simplicity
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return h.GetBooks(c) // Re-render the entire book list
}
