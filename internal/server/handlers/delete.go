package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a Api) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idB, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error mux.Vars: %w", err)
		return
	}

	id := int64(idB)
	ctx := r.Context()

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.storage.DeleteBook(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error encoder: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
