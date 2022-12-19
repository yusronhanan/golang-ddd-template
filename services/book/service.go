package book

import "golang-ddd-template/domain"

type Service struct {
	repo domain.BookRepository
}

func (s Service) AddBook(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateBook(book domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetBookByID(id int) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo domain.BookRepository) domain.BookService {
	return &Service{repo}
}
