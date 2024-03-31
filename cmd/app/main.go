package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	conn, err := pgx.Connect(context.Background(), "postgres://test:test@localhost:5432/test")
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

	http.ListenAndServe(":80", r)
}
