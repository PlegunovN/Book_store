package http2

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateRequest struct {
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func (a Api) CreateBook(W http.ResponseWriter, R *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	req := new(CreateRequest)
	_ = json.NewDecoder(R.Body).Decode(&req)

	ctx := context.TODO()
	err := a.Storage.Insert(ctx, req.Title, req.Author.Firstname, req.Author.Lastname)
	// обработать ошибку
	// записать ошибку в респонс
	if err != nil {
		log.Fatal("err in create.go")
	}

}
