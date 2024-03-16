package http2

import (
	"Book_store/internal/database"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	Storage *database.Service
}

func New(db *sqlx.DB) *Api {
	return &Api{
		Storage: database.New(db),
	}
}

//func New(db *sqlx.DB) *Service {
//	return &Service{
//		client: &client{
//			db: db,
//		},
//	}
//}
