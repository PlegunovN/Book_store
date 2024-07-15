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
		a.SLogger.Info("error decode update one author in updateauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("error, not id in request")
		return
	}

	if req.Firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("error, not first name in request")
		return
	}

	if req.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("error, not last name in request")
		return
	}

	err = a.Storage.UpdateAuthor(ctx, req.Firstname, req.Lastname, req.ID)
	if err != nil {
		a.SLogger.Info("error update one author in updateauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
