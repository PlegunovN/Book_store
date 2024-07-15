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
		a.SLogger.Info("error mux.Vars book in getbook.go")
		return
	}

	ctx := context.WithValue(context.Background(), r, idStr)
	id := int64(idStr)

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("error, not ID")
		return
	}

	book, err := a.Storage.SelectBook(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Info("error select book")
		return
	}

	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		a.SLogger.Info("error, Book Not Found")
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Info("error Encode author in getbook.go")
		return
	}
	w.WriteHeader(http.StatusOK)
}
