package server

import (
	"fmt"
	"github.com/PlegunovN/Book_store/internal/books"
	"github.com/PlegunovN/Book_store/internal/server/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func ServerStart(storage *books.Service, sLogger *zap.SugaredLogger) {

	api := handlers.Api{Storage: storage,
		SLogger: sLogger}

	r := mux.NewRouter()
	fmt.Println("server start at 8080")
	sLogger.Info("hi logger")

	r.HandleFunc("/books", api.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", api.GetBook).Methods("GET")
	r.HandleFunc("/author/{id}", api.GetAuthor).Methods("Get")
	r.HandleFunc("/book", api.CreateBook).Methods("POST")
	r.HandleFunc("/update/book_author", api.UpdateBookandAuthor).Methods("PUT")
	r.HandleFunc("/update/book", api.UpdateBook).Methods("PUT")
	r.HandleFunc("/update/author", api.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/book/{id}", api.DeleteBook).Methods("DELETE")
	err := http.ListenAndServe(":8080", r)
	sLogger.Fatal(err)
}
