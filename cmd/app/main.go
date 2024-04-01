package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	conn, err := pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@localhost:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB")))
	if err != nil {
		log.Error(err)
	}
	defer conn.Close(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := conn.Ping(r.Context()); err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("welcome"))
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r)
}
