package handlers

import (
	"Book_store/internal/books"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	Storage *books.Service
}

func New(db *sqlx.DB) *Api {
	return &Api{
		Storage: books.New(db),
	}
}
