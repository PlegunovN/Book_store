package main

import (
	"Book_store/internal/books"
	"Book_store/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// авторизация
// логер

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
	if err != nil {
		log.Fatal("not connected to db")
	}
	storage := books.New(db)
	server.ServerStart(storage)

}
