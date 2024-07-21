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
		return
	}

	offset := r.URL.Query().Get("offset")
	if offset == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	books, err := a.Storage.SelectBooks(ctx, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorln("error select books")
		return
	}

	if books == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorln("error Encode books in getbooks.go")
		return
	}
	w.WriteHeader(http.StatusOK)
}
