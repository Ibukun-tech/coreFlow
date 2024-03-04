package main

import (
	"net/http"

	"github.com/Ibukun-tech/coreFlow/pkg/Handler"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// mux := pat.New()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Post("/available-json", http.HandlerFunc(Handler.Repo.AvailableJson))
	mux.Post("/search-availability", http.HandlerFunc(Handler.Repo.PostAvailability))
	mux.Get("/", http.HandlerFunc(Handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(Handler.Repo.About()))
	mux.Post("/make-reservation", http.HandlerFunc(Handler.Repo.PostReservation))
	return mux
}
