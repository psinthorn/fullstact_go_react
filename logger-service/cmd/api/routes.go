package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	// create new router
	mux := chi.NewRouter()
	// enable cors with handler
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, //3 seconds
	}))

	// Test route with middelware
	mux.Use(middleware.Heartbeat("/ping"))

	//  Router link list start here
	mux.Get("/welcome", app.Welcome)
	mux.Post("/logger", app.NewLog)

	return mux
}
