package handlers

import (
	"encoding/json"
	"net/http"
)

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateRequest struct {
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func (a Api) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(CreateRequest)
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error in decoder: %w", err)
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Author.Firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Author.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.storage.Insert(ctx, req.Title, req.Author.Firstname, req.Author.Lastname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("err encoder: %w", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
