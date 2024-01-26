package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stellyes/go-bed-n-breakfast/pkg/config"
	"github.com/stellyes/go-bed-n-breakfast/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	// Creates http handler
	mux := chi.NewRouter()

	// Handles unexpected app crashing, built-in middlware to Chi
	mux.Use(middleware.Recoverer)

	// Homemade middleware from middleware.go!
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Create http routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
