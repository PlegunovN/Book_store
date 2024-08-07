package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type UpdateAuthor struct {
	ID        int64  `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}

func (a Api) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(UpdateAuthor)
	ctx := context.WithValue(context.Background(), req, UpdateAuthor{})

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error decode update one author: %w", err)
		return
	}

	if req.ID == 0 {
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

	err = a.storage.UpdateAuthor(ctx, req.Firstname, req.Lastname, req.ID)
	if err != nil {
		a.logger.Errorf("error update one author: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
