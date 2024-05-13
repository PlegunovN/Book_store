package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type UpdateBookAuthor struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}

type UpdateBook struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type UpdateAuthor struct {
	ID        int64  `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}

func (a Api) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(UpdateBook)
	ctx := context.WithValue(context.Background(), req, UpdateBook{})
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error decode update one book in update.go")
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = a.Storage.UpdateBook(ctx, req.Title, req.ID)
	defer w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println("error update one book in update.go")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a Api) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(UpdateAuthor)
	ctx := context.WithValue(context.Background(), req, UpdateAuthor{})
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error decode update one author in update.go")
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = a.Storage.UpdateAuthor(ctx, req.Firstname, req.Lastname, req.ID)
	defer w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println("error update one author in update.go")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a Api) UpdateBookandAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(UpdateBookAuthor)
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error decode update book&author in update.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not ID")
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not Title")
		return
	}

	if req.Firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not Fist name")
		return
	}

	if req.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not Last name")
		return
	}

	err = a.Storage.UpdateBookAndAuthor(ctx, req.Title, req.ID, req.Firstname, req.Lastname)
	if err != nil {
		log.Println("error update book&author in update.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
