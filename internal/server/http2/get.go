package http2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// 29.02
func (a Api) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	ctx := context.TODO()
	books, err := a.Storage.SelectBooks(ctx, limit, offset)
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		fmt.Println("error Encode books in get.go")
	}
}

// 28,02
func (a Api) GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()
	idA, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("error mux.Vars book in get.go")
	}
	id := int64(idA)
	author, err1 := a.Storage.SelectAuthor(ctx, id)
	err1 = json.NewEncoder(w).Encode(author)
	if err1 != nil {
		fmt.Println("error Encode author in get.go")
	}
}

// 26,02
func (a Api) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()
	idB, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("error mux.Vars book in get.go")
	}
	id := int64(idB)
	book, err := a.Storage.SelectBook(ctx, id)
	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		fmt.Println("error Encode book in get.go")
	}
}

// 25,02
//func (a Api) GetBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Get("Content-Type")
//	ctx := context.TODO()
//
//	idB := r.URL.Query().Get("id")
//	id, err := strconv.ParseInt(idB, 10, 64)
//	if err != nil {
//		fmt.Println("error ParseInt in get.go")
//		fmt.Println("id", id)
//		fmt.Println("idSTR", idB)
//	}
//	book, err1 := a.Storage.SelectBook(ctx, id)
//
//	err1 = json.NewEncoder(w).Encode(book)
//	if err1 != nil {
//		return
//	}

// принять ошибку
//if err1 != nil {
//	log.Fatal("err in get.go")
//}

//20,02
//func (a Api) GetBooks(W http.ResponseWriter, R *http.Request) {
//	W.Header().Set("Content-Type", "application/json")
//
//	ctx := context.TODO()
//
//	books, err := a.Storage.SelectAll(ctx)
//	if err != nil {
//		log.Fatal("err in create.go")
//
//	}
//	json.NewEncoder(W).Encode(books)
//
//}

//17,02
//func (a Api) GetBooks(W http.ResponseWriter, R *http.Request) Book {
//	W.Header().Set("Content-Type", "application/json")
//ctx := context.TODO()
//req := new(CreateRequest)
//err := a.Storage.Insert(ctx, req.Title, req.Author.Firstname, req.Author.Lastname)
//if err != nil {
//	log.Fatal("err in create.go")
//books := a.Storage
//json.NewEncoder(W).Encode(books)
//return books
//}

//old
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	books := db.GetBooks()
//	json.NewEncoder(w).Encode(books)
//}
//
///*
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(db.books)
//}
//*/
//
//func GetBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	id := r.URL.Query().Get("id")
//	book := db.GetBook(id)
//	json.NewEncoder(w).Encode(book)
//}
