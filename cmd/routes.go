package cmd

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"shaw-and-partners-technical-challenge/controllers"
	"shaw-and-partners-technical-challenge/middlewares"
)

func (app *Application) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.SetAcceptMiddleware)

	mux.Get("/", controllers.Home)
	mux.Get("/api/users", controllers.GetAllUsers)
	mux.Get("/api/users/{username}/details", controllers.GetUserDetails)
	mux.Get("/api/users/{username}/repos", controllers.GetUserRepositories)

	fileServer := http.FileServer(http.Dir("./api/"))
	mux.Handle("/api/*", http.StripPrefix("/api", fileServer))
	return mux
}
