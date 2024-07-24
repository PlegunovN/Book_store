package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a Api) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error mux.Vars book in getbook.go: %w", err)
		return
	}

	ctx := context.WithValue(context.Background(), r, idStr)
	id := int64(idStr)

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := a.Storage.SelectBook(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error select book: %w", err)
		return
	}

	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error Encode author in getbook.go: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
