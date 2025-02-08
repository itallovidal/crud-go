package main

import (
	"log/slog"
	"net/http"
	"rocketseat-desafio-1/handlers"
	"rocketseat-desafio-1/middlewares"
	"rocketseat-desafio-1/models"
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

	db := map[string]models.User{
		"1703c5e5-d6e0-467f-840e-7b4b32ab3650":{
			First_name: "Itallo",
			Last_name: "Vidal",
			Biography: "Estudante de SI.",
		},
	}

	r.Route("/api", func(r chi.Router) {
		// idRegex:= "^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$}"

		r.Route("/users", func(r chi.Router) {
			r.Post("/", handlers.CreateUserController(db))
			r.Get("/", handlers.GetAllUsers(db))
			r.Get("/{id}", handlers.GetUserById(db))
			r.Delete("/{id}", handlers.DeleteUser(db))
			r.Put("/{id}", handlers.UpdateUser(db))
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