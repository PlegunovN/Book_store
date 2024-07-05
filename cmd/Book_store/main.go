package main

import (
	"Book_store/internal"
	"Book_store/internal/books"
	"Book_store/internal/server"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// авторизация
// логер

func main() {
	internal.InitLogger()
	defer internal.SugarLogger.Sync()

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
	if err != nil {
		internal.SugarLogger.Fatal("not connected to db")
		fmt.Println("not connected to db")
	}
	storage := books.New(db)
	server.ServerStart(storage)

}
