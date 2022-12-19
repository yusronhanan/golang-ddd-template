package domain

import (
	"github.com/labstack/echo/v4"
	"time"
)

type (
	Book struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Owner     string    `json:"owner"`
		CreatedAt time.Time `json:"created_at"`
	}

	BookHandler interface {
		AddBook(c echo.Context) error
		UpdateBook(c echo.Context) error
		GetBookByID(c echo.Context) error
	}
	BookService interface {
		AddBook(book Book) error
		UpdateBook(book Book) error
		GetBookByID(id int) (*Book, error)
	}

	BookRepository interface {
		Insert(book Book) error
		Update(book Book) error
		GetByID(id int) (*Book, error)
	}
)
