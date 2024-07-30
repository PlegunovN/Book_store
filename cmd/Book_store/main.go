package main

import (
	"Book_store/internal/books"
	"Book_store/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// авторизация
// логер

func main() {
	//config init
	cfg, err := configs.LoadConfig("./.env")
	if err != nil {
		log.Fatal(err)
	}


	sLogger := logger.InitLogger()
	defer sLogger.Sync()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode))
	if err != nil {
		sLogger.Fatalf("not connected to db: %w", err)
	}
	storage := books.New(db, sLogger)
	server.ServerStart(storage, sLogger)

}
