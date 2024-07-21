package handlers

import (
	"context"
	"encoding/json"
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
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorln("error decode update one book in updatebook.go")
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

	err = a.Storage.UpdateBook(ctx, req.Title, req.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorln("error update one book in updatebook.go")
		return
	}

	w.WriteHeader(http.StatusOK)
}
