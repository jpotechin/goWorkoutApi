package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jpotechin/goWorkoutApi/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	return r
}
