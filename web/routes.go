package main

import (
	"booking/internal/config"
	"booking/internal/handler"
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
	//mux.Get("/search-availability-json", handler.Repo.AvailabilityJSON)
	mux.Post("/search-availability-json", handler.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handler.Repo.ChooseRoom)
	mux.Get("/book-room", handler.Repo.BookRoom)

	mux.Get("/contact", handler.Repo.Contact)

	mux.Get("/make-reservation", handler.Repo.Reservation)
	mux.Post("/make-reservation", handler.Repo.PostReservation)
	mux.Get("/reservation-summery", handler.Repo.ReservationSummery)

	mux.Get("/user/login", handler.Repo.ShowLogin)
	mux.Post("/user/login", handler.Repo.PostShowLogin)
	mux.Get("/user/logout", handler.Repo.Logout)

	fileServer := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handler.Repo.AdminDashboard)

		mux.Get("/reservations-new", handler.Repo.AdminNewReservation)
		mux.Get("/reservations-all", handler.Repo.AdminAllReservation)
		mux.Get("/reservations-calendar", handler.Repo.AdminReservationCalendar)
		mux.Post("/reservations-calendar", handler.Repo.AdminPostReservationsCalendar)

		mux.Get("/reservations/{src}/{id}/show", handler.Repo.AdminShowReservation)
		mux.Post("/reservations/{src}/{id}", handler.Repo.AdminPostShowReservation)

		mux.Get("/process-reservation/{src}/{id}/do", handler.Repo.AdminProcessReservation)
		mux.Get("/delete-reservation/{src}/{id}/do", handler.Repo.AdminDeleteReservation)
	})
	return mux
}
