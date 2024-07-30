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
		a.logger.Errorf("error mux.Vars : %w", err)
		return
	}
	ctx := r.Context()
	id := int64(idStr)

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author, err := a.storage.SelectAuthor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error select author: %w", err)
		return
	}
	if author == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error encoder: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
