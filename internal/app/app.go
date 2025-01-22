package app

import (
	"database/sql"
	"github.com/go_web/internal/configs"
	"github.com/go_web/internal/delivery/http/handlers"
	"github.com/go_web/internal/repository"
	"github.com/go_web/internal/service"
	"log"
	"net/http"
	"time"
)

func RunServer(cfg configs.Config) {
	db := repository.InitDB(cfg.Postgres)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Not closer DB")
		}
	}(db)

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)

	srv := http.Server{
		Addr:           ":8080",
		Handler:        handlers.RegisterRoutes(handler),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
