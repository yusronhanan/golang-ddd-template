package notes

import (
	"github.com/labstack/echo/v4"
	"golang-ddd-template/domain"
	"golang-ddd-template/public"
)

type Handler struct {
	svc domain.NotesService
}

func (h Handler) AddNotes(c echo.Context) error {
	var request public.ReqNotes
	_ = request
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateNotes(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetNotesByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetNotesByBookID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewHandler(svc domain.NotesService) domain.NotesHandler {
	return &Handler{svc}
}
