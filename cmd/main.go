package cmd

import (
	"github.com/labstack/echo/v4"
	"golang-ddd-template/config"
	"golang-ddd-template/domain"
	"golang-ddd-template/services/book"
	"golang-ddd-template/services/notes"
	"log"
)

func main() {
	env := config.ENV()
	_ = env

	repoBook := book.NewRepositoryMySql("")
	svcBook := book.NewService(repoBook)
	handlerBook := book.NewHandler(svcBook)

	repoNotes := notes.NewRepository("")
	svcNotes := notes.NewService(repoNotes)
	handlerNotes := notes.NewHandler(svcNotes)
	Routes(env.AppPort, handlerNotes, handlerBook)
}

func Routes(appPort string, handlerNotes domain.NotesHandler, handlerBook domain.BookHandler) {
	e := echo.New()
	//define route here
	log.Println(e.Start("0.0.0.0:" + appPort))
}
