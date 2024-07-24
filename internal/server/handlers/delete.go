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
		a.SLogger.Errorf("error mux.Vars book in delete.go: %w", err)
		return
	}

	id := int64(idB)
	ctx := r.Context()

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.Storage.DeleteBook(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error Encode id in delete.go: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
