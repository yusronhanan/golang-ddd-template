package notes

import "golang-ddd-template/domain"

type Repository struct {
	db string
}

func (r Repository) Insert(Notes domain.Notes) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(Notes domain.Notes) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetByID(id int) (*domain.Notes, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetByBookID(bookId int) ([]domain.Notes, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db string) domain.NotesRepository {
	return &Repository{db}
}
