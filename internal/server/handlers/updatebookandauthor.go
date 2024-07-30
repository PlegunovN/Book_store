package handlers

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error decode update book&author : %w", err)
		return
	}

	if req.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.storage.UpdateBookAndAuthor(ctx, req.Title, req.ID, req.Firstname, req.Lastname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error update book&author : %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
