package server

import (
	"Book_store/internal/books"
	"Book_store/internal/logger"
	"Book_store/internal/server/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ServerStart(storage *books.Service) {

	api := handlers.Api{Storage: storage}

	r := mux.NewRouter()
	fmt.Println("server start at 8080")
	logger.SugarLogger.Info("hi logger")

	r.HandleFunc("/books", api.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", api.GetBook).Methods("GET")
	r.HandleFunc("/author/{id}", api.GetAuthor).Methods("Get")
	r.HandleFunc("/book", api.CreateBook).Methods("POST")
	r.HandleFunc("/update/book_author", api.UpdateBookandAuthor).Methods("PUT")
	r.HandleFunc("/update/book", api.UpdateBook).Methods("PUT")
	r.HandleFunc("/update/author", api.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/book/{id}", api.DeleteBook).Methods("DELETE")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
