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
	//"Book_store/internal/server/http2"

	//"Book_store/internal/database"
	"Book_store/internal/database"
	"Book_store/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
	if err != nil {
		log.Fatal("not connected to db")
	}
	storage := database.New(db)
	server.ServerStart(storage)

}
