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
	logger.InitLogger()
	defer logger.SugarLogger.Sync()

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
	if err != nil {
		logger.SugarLogger.Fatal("not connected to db")
		fmt.Println("not connected to db")
	}
	storage := books.New(db, logger.SugarLogger)
	server.ServerStart(storage)

}
