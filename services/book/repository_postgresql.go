package book

import "golang-ddd-template/domain"

type RepositoryPostgresql struct {
	db string
}

func (r RepositoryPostgresql) Insert(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryPostgresql) Update(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryPostgresql) GetByID(id int) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepositoryPostgresql(db string) domain.BookRepository {
	return &RepositoryPostgresql{db}
}
