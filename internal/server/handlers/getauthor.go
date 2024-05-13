package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (a Api) GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error mux.Vars book in getauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	id := int64(idStr)

	author, err := a.Storage.SelectAuthor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("error Not Found")
		return
	}

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error, not ID")
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		log.Println("error Encode author in getauthor.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
