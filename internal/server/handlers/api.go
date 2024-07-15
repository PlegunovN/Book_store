package handlers

import (
	"Book_store/internal/books"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Api struct {
	Storage *books.Service
}

func New(db *sqlx.DB, logger *zap.SugaredLogger) *Api {
	return &Api{
		Storage: books.New(db, logger),
	}
}
