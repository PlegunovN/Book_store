package main

import (
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
		sLogger.Fatalf("not connected to db: %w", err)
	}
	storage := books.New(db, sLogger)
	server.ServerStart(storage, sLogger)

}
