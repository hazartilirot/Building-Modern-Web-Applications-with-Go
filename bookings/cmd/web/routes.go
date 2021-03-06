package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/internal/config"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/majors-suite", handlers.Repo.MajorsSuite)
	mux.Get("/generals-quarters", handlers.Repo.GeneralsQuarters)
	mux.Get("/room-availability", handlers.Repo.RoomAvailability)
	mux.Post("/room-availability", handlers.Repo.FormRoomAvailability)
	mux.Post("/room-availability-json", handlers.Repo.FormRoomAvailabilityJson)
	mux.Get("/reservations", handlers.Repo.Reservations)
	mux.Post("/reservations", handlers.Repo.PostReservations)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
