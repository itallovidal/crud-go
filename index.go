package main

import (
	"log/slog"
	"net/http"
	"rocketseat-desafio-1/controllers"
	"rocketseat-desafio-1/interfaces"
	"rocketseat-desafio-1/middlewares"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main(){
	slog.Info("Server Initializing..")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middlewares.SetJsonResponseMiddleware)

	db := map[string]interfaces.User{}

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", controllers.CreateUserController(db))
			// r.Get("/")
			// r.Get("/{id}")
			// r.Delete("/{id}")
			// r.Put("/{id}")
		})
	})

	app := http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("\nServer initialized.\n")
	serverErr := app.ListenAndServe()

	if serverErr != nil {
		slog.Error("Something went wrong, server is closed.")
	} 
}