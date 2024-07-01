package main

import (
	"Book_store/configs"
	"Book_store/internal/books"
	"Book_store/internal/server"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// авторизация
// логер

func main() {
	//config init
	cfg, err := configs.LoadConfig("./.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode))
	if err != nil {
		log.Fatal("not connected to db")
	}
	storage := books.New(db)
	server.ServerStart(storage)
	//defer db.Close()
}
