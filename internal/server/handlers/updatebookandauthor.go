package handlers

import (
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

func (a Api) UpdateBookandAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(UpdateBookAuthor)
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error decode update book&author in updatebookandauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not IDin request")
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not BookTitle in request")
		return
	}

	if req.Firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not Fist name in request")
		return
	}

	if req.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not Last name in request")
		return
	}

	err = a.Storage.UpdateBookAndAuthor(ctx, req.Title, req.ID, req.Firstname, req.Lastname)
	if err != nil {
		log.Println("error update book&author in updatebookandauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
