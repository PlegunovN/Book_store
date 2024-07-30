package handlers

import (
	"github.com/PlegunovN/Book_store/internal/books"
	"go.uber.org/zap"
)

type Api struct {
	storage *books.Service
	logger  *zap.SugaredLogger
}

func New(storage *books.Service, logger *zap.SugaredLogger) *Api {
	return &Api{
		storage: storage,
		logger:  logger,
	}
}
