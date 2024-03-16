package http2

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// 28,02
func (a Api) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()
	idB, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("error mux.Vars book in delete.go")
	}
	id := int64(idB)
	err = a.Storage.DeleteBook(ctx, id)
	if err != nil {
		fmt.Println("error Encode book in delete.go")
	}
}

//old
//func DeleteBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	id := r.URL.Query().Get("id")
//	db.DeleteBook(id)
//	books := db.GetBooks()
//	json.NewEncoder(w).Encode(books)
//}
