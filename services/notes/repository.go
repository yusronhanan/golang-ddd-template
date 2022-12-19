package notes

import "golang-ddd-template/domain"

type Repository struct {
	db string
}

func (r Repository) Insert(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetByID(id int) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db string) domain.BookRepository {
	return &Repository{db}
}
