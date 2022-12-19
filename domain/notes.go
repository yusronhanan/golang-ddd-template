package domain

import (
	"github.com/labstack/echo/v4"
	"time"
)

type (
	Notes struct {
		ID        int       `json:"id"`
		BookID    int       `json:"book_id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
	}

	NotesHandler interface {
		AddNotes(c echo.Context) error
		UpdateNotes(c echo.Context) error
		GetNotesByID(c echo.Context) error
		GetNotesByBookID(c echo.Context) error
	}
	NotesService interface {
		AddNotes(Notes Notes) error
		UpdateNotes(Notes Notes) error
		GetNotesByID(id int) (*Notes, error)
		GetNotesByBookID(bookId int) ([]Notes, error)
	}

	NotesRepository interface {
		Insert(Notes Notes) error
		Update(Notes Notes) error
		GetByID(id int) (*Notes, error)
		GetByBookID(bookId int) ([]Notes, error)
	}
)
