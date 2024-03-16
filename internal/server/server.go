package server

import (
	"Book_store/internal/database"
	"Book_store/internal/server/http2"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerStart(storage *database.Service) {

	api := http2.Api{Storage: storage}

	r := mux.NewRouter()
	fmt.Println("server start at 8080")
	r.HandleFunc("/books", api.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", api.GetBook).Methods("GET")
	r.HandleFunc("/author/{id}", api.GetAuthor).Methods("Get")
	r.HandleFunc("/book", api.CreateBook).Methods("POST")
	r.HandleFunc("/update/book", api.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", api.DeleteBook).Methods("DELETE")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
