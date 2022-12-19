package book

import (
	"github.com/labstack/echo/v4"
	"golang-ddd-template/domain"
	"golang-ddd-template/public"
)

type Handler struct {
	svc domain.BookService
}

func (h Handler) AddBook(c echo.Context) error {
	var request public.ReqBook
	_ = request
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateBook(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetBookByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewHandler(svc domain.BookService) domain.BookHandler {
	return &Handler{svc}
}
