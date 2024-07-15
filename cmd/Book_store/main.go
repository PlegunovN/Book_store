package main

import (
	"fmt"
	"github.com/PlegunovN/Book_store/internal/books"
	"github.com/PlegunovN/Book_store/internal/logger"
	"github.com/PlegunovN/Book_store/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// авторизация
// логер

func main() {
	sLogger := logger.InitLogger()
	defer sLogger.Sync()

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
	if err != nil {
		fmt.Println("not connected to db")
		sLogger.Fatal("not connected to db")
	}
	storage := books.New(db, sLogger)
	server.ServerStart(storage, sLogger)

}
