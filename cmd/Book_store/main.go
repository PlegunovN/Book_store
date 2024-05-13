package main

import (
	//"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	//"log"
	//"math/rand"
	//"net/http"
	//"Book_store/internal/services"
	//"strconv"
	//"Book_store/internal/server/handlers"

	//"Book_store/internal/books"
	"Book_store/internal/books"
	"Book_store/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// миграции
// конфиг
// /авторизация
// логер
// сделать связь книги автор многие ко многим
func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
	if err != nil {
		log.Fatal("not connected to db")
	}
	storage := books.New(db)
	server.ServerStart(storage)

}

//create table author
//(
//id        bigserial
//primary key,
//firstname varchar(50),
//lastname  varchar(50)
//);
//create table book
//(
//id     bigserial
//primary key,
//title  varchar(50),
//author bigint
//references author
//);
