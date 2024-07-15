package handlers

import (
	"github.com/PlegunovN/Book_store/internal/books"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Api struct {
	Storage *books.Service
	SLogger *zap.SugaredLogger
}

func New(db *sqlx.DB, logger *zap.SugaredLogger) *Api {
	return &Api{
		Storage: books.New(db, logger),
	}
}
