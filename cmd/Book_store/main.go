package main

import (
	"Book_store/internal/books"
	"Book_store/internal/server"
	_ "github.com/lib/pq"
)

// авторизация
// логер

func main() {
	//config init

	storage := books.New(db)
	server.ServerStart(storage)
	//defer db.Close()
}
