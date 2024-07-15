package handlers

import (
	"Book_store/internal/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a Api) GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.SugarLogger.Info("error mux.Vars book in getauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	id := int64(idStr)

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.SugarLogger.Info("error, not ID in request")
		return
	}

	author, err := a.Storage.SelectAuthor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.SugarLogger.Info("error select author")
		return
	}
	if author == nil {
		w.WriteHeader(http.StatusNotFound)
		logger.SugarLogger.Info("error, Author Not Found")
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		logger.SugarLogger.Info("error Encode author in getauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
