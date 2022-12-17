package main

import (
	"net/http"

	"github.com/Tedrick10/go-bookings/pkg/config"
	"github.com/Tedrick10/go-bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
);

func routes(app *config.AppConfig) http.Handler {
	/* Example with pat */
	// mux := pat.New();
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home));
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About));

	/* Example with chi */
	mux := chi.NewRouter();
	mux.Use(middleware.Recoverer);
	mux.Use(SessionLoad);

	// mux.Use(WriteToConsole);
	mux.Use(NoSurf);
	mux.Get("/", handlers.Repo.Home);
	mux.Get("/about", handlers.Repo.About);

	fileServer := http.FileServer(http.Dir("./static"));
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer));
	
	return mux;
}