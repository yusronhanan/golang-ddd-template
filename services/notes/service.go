package notes

import "golang-ddd-template/domain"

type Service struct {
	repo domain.NotesRepository
}

func (s Service) AddNotes(Notes domain.Notes) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateNotes(Notes domain.Notes) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetNotesByID(id int) (*domain.Notes, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetNotesByBookID(bookId int) ([]domain.Notes, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo domain.NotesRepository) domain.NotesService {
	return &Service{repo}
}
