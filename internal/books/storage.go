package books

import (
	"Book_store/internal/logger"
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type client struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func (s client) insert(ctx context.Context, book Book, author Author) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.SugarLogger.Info("Create error - rollback")
		}
		tx.Commit()
	}()

	query := "INSERT INTO author(firstname, lastname) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, author.Firstname, author.Lastname)
	if err != nil {
		return err
	}

	query = "INSERT INTO book(title) VALUES($1)"
	_, err = tx.ExecContext(ctx, query, book.Title)

	query = "SELECT id FROM author where  firstname = $1 AND lastname = $2"
	a := new(Author)
	err = tx.GetContext(ctx, a, query, author.Firstname, author.Lastname)
	if err != nil {
		return err
	}

	query = "SELECT id FROM book WHERE title = $1"
	b := new(Book)
	err = tx.GetContext(ctx, b, query, book.Title)
	if err != nil {
		return err
	}

	query = "INSERT INTO bookauthor(author_id, book_id) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, query, a.ID, b.ID)

	return err
}

func (s client) UpdateBookAndAuthor(ctx context.Context, title string, id int64, firstname, lastname string) error {

	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.SugarLogger.Info("update error - rollback ")
			return
		}
		tx.Commit()
	}()

	_, err = tx.ExecContext(ctx, "UPDATE book SET title = $1 WHERE id = $2", title, id)
	if err != nil {
		return err
	}
	b := new(Book)
	query := "SELECT author.id FROM author INNER JOIN bookAuthor ON author.id = bookAuthor.author_id " +
		"INNER JOIN book on bookAuthor.book_id = book.id WHERE book.id = $1"
	err = tx.GetContext(ctx, b, query, id)
	if err != nil {
		return err
	}

	query = "UPDATE author SET firstname = $2, lastname = $3 WHERE id=$1"
	_, err = tx.ExecContext(ctx, query, b.ID, firstname, lastname)
	if err != nil {
		return err
	}

	return nil
}

func (s client) UpdateBook(ctx context.Context, title string, id int64) error {
	query := "UPDATE book SET title = $1 WHERE id = $2"
	_, err := s.db.ExecContext(ctx, query, title, id)
	if err != nil {
		return err
	}
	return nil
}

func (s client) UpdateAuthor(ctx context.Context, firstname, lastname string, id int64) error {
	query := "UPDATE author SET firstname = $1, lastname = $2 WHERE id = $3"
	_, err := s.db.ExecContext(ctx, query, firstname, lastname, id)
	if err != nil {
		return err
	}
	return nil
}

func (s client) GetBooks(ctx context.Context, limit, offset string) ([]BookAuthor, error) {
	query := `"SELECT book.id, book.title, author.id , author.firstname, author.lastname 
		FROM book INNER JOIN bookAuthor ON book.id = bookAuthor.book_id 
		INNER JOIN author ON bookAuthor.author_id = author.id  ORDER BY book.id LIMIT $1 OFFSET $2"`

	books := make([]BookAuthor, 1)
	err := s.db.SelectContext(ctx, &books, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s client) GetBook(ctx context.Context, id int64) (*BookAuthor, error) {
	query := `"SELECT book.id, book.title , author.id, author.firstname, author.lastname
		FROM book INNER JOIN bookAuthor ON book.id = bookAuthor.book_id 
		INNER JOIN author ON bookAuthor.author_id = author.id where book.id=$1"`
	res := &BookAuthor{}
	err := s.db.GetContext(ctx, res, query, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s client) GetAuthor(ctx context.Context, id int64) (*Author, error) {
	query := "SELECT id, firstname, lastname FROM author WHERE id=$1"
	author := &Author{}
	err := s.db.GetContext(ctx, author, query, id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s client) DeleteBook(ctx context.Context, id int64) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.SugarLogger.Info("Delete error - rollback")
		}
		tx.Commit()
	}()
	query := "DELETE FROM bookauthor WHERE book_id = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	query = "DELETE FROM book WHERE book.id=$1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
