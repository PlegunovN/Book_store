package handlers

import (
	"context"
	"encoding/json"
	"github.com/PlegunovN/Book_store/internal/logger"
	"net/http"
)

type UpdateBook struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (a Api) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(UpdateBook)
	ctx := context.WithValue(context.Background(), req, UpdateBook{})

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.SugarLogger.Info("error decode update one book in updatebook.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.SugarLogger.Info("error, not id in request")
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		logger.SugarLogger.Info("error, not title in request")
		return
	}

	err = a.Storage.UpdateBook(ctx, req.Title, req.ID)
	if err != nil {
		logger.SugarLogger.Info("error update one book in updatebook.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
