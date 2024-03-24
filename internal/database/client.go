package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type client struct {
	db *sqlx.DB
}

// $ - placeholder место куда подставятся значения
func (s client) insert(ctx context.Context, book Book, author Author) error {
	query := "INSERT INTO author (firstname, lastname) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, author.Firstname, author.Lastname)
	if err != nil {
		return err
	}

	query = "SELECT id, firstname, lastname FROM author WHERE firstname = $1 AND lastname = $2"

	a := new(Author)
	err = s.db.GetContext(ctx, a, query, author.Firstname, author.Lastname)
	if err != nil {
		return err
	}

	query = "INSERT INTO book(title, author) VALUES($1,$2)"
	_, err = s.db.ExecContext(ctx, query, book.Title, a.ID)
	return err
}

func (s client) UpdateBook(ctx context.Context, title string, id int64, firstname, lastname string) error {

	tx, err := s.db.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Println("update error - rollback ")
			return
		}
		tx.Commit()
	}()
	//Сделать отдельно обновление книга и автора
	//либо
	//добавить условия обновления
	_, err = tx.ExecContext(ctx, "UPDATE book SET title = $1 WHERE book.id = $2", title, id)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE author SET firstname = $3, lastname = $4 WHERE author.id = book.author", firstname, lastname)
	if err != nil {
		return err
	}

	return nil
}

func (s client) UpBook(ctx context.Context, title string, id int64) error {
	query := "UPDATE book SET title = $1 WHERE id = $2"
	_, err := s.db.ExecContext(ctx, query, title, id)
	if err != nil {
		return err
	}
	return nil
}

func (s client) UpAuthor(ctx context.Context, firstname, lastname string, id int64) error {
	query := "UPDATE author SET firstname = $1, lastname = $2 WHERE id = $3"
	_, err := s.db.ExecContext(ctx, query, firstname, lastname, id)
	if err != nil {
		return err
	}
	return nil
}

// 29,02
func (s client) GetBooks(ctx context.Context, limit, offset string) ([]BookAuthor, error) {
	query := "SELECT book.id, title, book.author, firstname, lastname  FROM book INNER JOIN author ON book.author = author.id  ORDER BY id LIMIT $1 OFFSET $2"
	books := make([]BookAuthor, 1)
	err := s.db.SelectContext(ctx, &books, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s client) GetBook(ctx context.Context, id int64) (*BookAuthor, error) {
	//query запрос на  книгу по ид
	query := "SELECT book.id, title, book.author, firstname, lastname  FROM book INNER JOIN author ON book.author = author.id where book.id=$1"
	//поинтер на книгу
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

// del book
func (s client) DelBook(ctx context.Context, id int64) error {
	query := "DELETE FROM book WHERE id=$1"

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

//func (S client) selectAll(ctx context.Context) ([]*Book, error) {
//
//	return nil, nil
//	//query := "SELECT * FROM author (fistname, lastname) VALUES ($1, $2)"
//	//author := make()
//	//err := S.db.GetContext(ctx, query, author.Firstname, author.Lastname)
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//query = "SELECT id, firstname, lastname FROM author WHERE firstname = $1 AND lastname = $2"
//	//
//	//a := new(Author)
//	//err = S.db.GetContext(ctx, a, query, author.Firstname, author.Lastname)
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//query = "SELECT * FROM book(title, author) VALUES($1, $2)"
//	//err = S.db.GetContext(ctx, query, book.Title, a.ID)
//	//return err
//}

//type db struct {
//	books []Book
//}

//func (d *db) SetBook(book Book) {
//	d.books = append(d.books, book)
//}
//
//func (d *db) DeleteBook(id int64) {
//	for index, item := range d.books {
//		if item.ID == id {
//			d.books = append(d.books[:index], d.books[index+1:]...)
//			break
//		}
//	}
//
//}
//
//func (d *db) GetBook(id int64) Book {
//	for _, item := range d.books {
//		if item.ID == id {
//			return item
//		}
//	}
//	return Book{}
//}
//
//func (d *db) GetBooks() []Book {
//	return d.books
//
//}
//
//func (d *db) UpdateBook(id int64) Book {
//	for index, item := range d.books {
//		if item.ID == id {
//			d.books = append(d.books[:index], d.books[index+1:]...)
//			var book Book
//			book.ID = id
//			d.books = append(d.books, book)
//			return Book{}
//		}
//	}
//	return Book{}
//}
//
//func NewDB() *db {
//	var d *db = &db{books: make([]Book, 0)}
//	d.SetBook(Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
//	d.SetBook(Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
//	return d
//}

//func Add(db *db)
//db.SetBook("1", "Война и Мир", &database.Author{Firstname: "Лев", Lastname: "Толстой"})
//db.SetBook("2", "Преступление и наказание", &database.Author{Firstname: "Фёдор", Lastname: "Достоевский"})
//}
//
