package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (a Api) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idB, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error mux.Vars book in delete.go")
		return
	}

	id := int64(idB)
	ctx := r.Context()

	err = a.Storage.DeleteBook(ctx, id)
	if err != nil {
		log.Println("error Encode id in delete.go")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
