package handlers

import (
	"github.com/PlegunovN/Book_store/internal/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a Api) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idB, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.SugarLogger.Info("error mux.Vars book in delete.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := int64(idB)
	ctx := r.Context()

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.SugarLogger.Info("error, not ID in request")
		return
	}

	err = a.Storage.DeleteBook(ctx, id)
	if err != nil {
		logger.SugarLogger.Info("error Encode id in delete.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
