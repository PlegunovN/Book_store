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
		//logger.SugarLogger.Info("error mux.Vars book in getauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	id := int64(idStr)

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		a.SLogger.Info("error, not ID in request")
		return
	}

	author, err := a.Storage.SelectAuthor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Info("error select author")
		return
	}
	if author == nil {
		w.WriteHeader(http.StatusNotFound)
		a.SLogger.Info("error, Author Not Found")
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		a.SLogger.Info("error Encode author in getauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
