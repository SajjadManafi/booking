package main

import (
	"booking/pkg/config"
	"booking/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//using chi
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/generals-quarters", handler.Repo.Generals)
	mux.Get("/majors-suite", handler.Repo.Majors)

	mux.Get("/search-availability", handler.Repo.Availability)
	mux.Post("/search-availability", handler.Repo.PostAvailability)
	mux.Get("/search-availability-json", handler.Repo.AvailabilityJSON)

	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/make-reservation", handler.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
