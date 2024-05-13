package books

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	client *client
}

// New функция инициальзирует объект , работающий с бд (вид функций -creator)
func New(db *sqlx.DB) *Service {
	return &Service{
		client: &client{
			db: db,
		},
	}
}

// создание новой книги
func (s Service) Insert(ctx context.Context, title, authorFirstname, authorLastname string) error {
	err := s.client.insert(ctx, Book{BookTitle: title}, Author{AuthorFirstname: authorFirstname, AuthorLastname: authorLastname})
	return err
}

// получение всех книг
func (s Service) SelectBooks(ctx context.Context, limit, offset string) ([]BookAuthor, error) {
	books, err := s.client.GetBooks(ctx, limit, offset)
	return books, err
}

// получение одной книги по айди
func (s Service) SelectBook(ctx context.Context, id int64) (*BookAuthor, error) {
	book, err := s.client.GetBook(ctx, id)
	return book, err
}

// получение одного автора по айди
func (s Service) SelectAuthor(ctx context.Context, id int64) (*Author, error) {
	author, err := s.client.GetAuthor(ctx, id)
	return author, err
}

// удаление одной книги
func (s Service) DeleteBook(ctx context.Context, id int64) error {
	err := s.client.DeleteBook(ctx, id)
	return err
}

// update one book and authors
func (s Service) UpdateBookAndAuthor(ctx context.Context, title string, id int64, firstname, lastname string) error {
	err := s.client.UpdateBookAndAuthor(ctx, title, id, firstname, lastname)
	return err
}

// update one book
func (s Service) UpdateBook(ctx context.Context, title string, id int64) error {
	err := s.client.UpdateBook(ctx, title, id)
	return err
}

// update one author
func (s Service) UpdateAuthor(ctx context.Context, firstname, lastname string, id int64) error {
	err := s.client.UpdateAuthor(ctx, firstname, lastname, id)
	return err
}
