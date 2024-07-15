package handlers

import (
	"encoding/json"
	"net/http"
)

func (a Api) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("400", "error, not limit in request")
		return
	}

	offset := r.URL.Query().Get("offset")
	if offset == "" {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("error, not offset in request")
		return
	}

	books, err := a.Storage.SelectBooks(ctx, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Info("error select books")
		return
	}

	if books == nil {
		w.WriteHeader(http.StatusNotFound)
		a.SLogger.Info("error, books Not Found")
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Info("error Encode books in getbooks.go")
		return
	}
	w.WriteHeader(http.StatusOK)
}
