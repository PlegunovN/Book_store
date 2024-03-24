package http2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateB struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}

type UBook struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type UAuthor struct {
	ID        int64  `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}

// 23,03
func (a Api) UpBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()
	req := new(UBook)
	_ = json.NewDecoder(r.Body).Decode(&req)
	err := a.Storage.UpBook(ctx, req.Title, req.ID)
	if err != nil {
		fmt.Println(" error encode update one book in update.go")
	}
}

func (a Api) UpAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()
	req := new(UAuthor)
	_ = json.NewDecoder(r.Body).Decode(&req)
	err := a.Storage.UpAuthor(ctx, req.Firstname, req.Lastname, req.ID)
	if err != nil {
		fmt.Println(" error encode update one book in update.go")
	}
}

// 05,03
func (a Api) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	req := new(UpdateB)
	_ = json.NewDecoder(r.Body).Decode(&req)
	//idB, err := strconv.Atoi(mux.Vars(r)["id"])
	//if err != nil {
	//	fmt.Println("error mux.Vars book in update.go")
	//}
	//id := int64(idB)
	err := a.Storage.UpdateBook(ctx, req.Title, req.ID, req.Firstname, req.Lastname)
	if err != nil {
		fmt.Println("error Encode bookAuthor in update.go")
	}
}

//func UpdateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	id := r.URL.Query().Get("id")
//	book := db.UpdateBook(id)
//	json.NewEncoder(w).Encode(book)
//
//}

/*
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			var book database.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
*/
