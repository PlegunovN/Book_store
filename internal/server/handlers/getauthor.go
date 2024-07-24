package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a Api) GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error mux.Vars book in getauthor.go: %w", err)
		return
	}
	ctx := r.Context()
	id := int64(idStr)

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author, err := a.Storage.SelectAuthor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error select author: %w", err)
		return
	}
	if author == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error Encode author in getauthor.go: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
