package main

import (
	"net/http"

	"github.com/Ibukun-tech/coreFlow/pkg/Handler"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(Handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(Handler.Repo.About()))
	return mux
}
