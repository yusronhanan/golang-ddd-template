package book

import "golang-ddd-template/domain"

type RepositoryMySql struct {
	db string
}

func (r RepositoryMySql) Insert(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryMySql) Update(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryMySql) GetByID(id int) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepositoryMySql(db string) domain.BookRepository {
	return &RepositoryMySql{db}
}
