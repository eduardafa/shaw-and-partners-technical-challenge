package cmd

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"shaw-and-partners-technical-challenge/controllers"
	"shaw-and-partners-technical-challenge/middlewares"
)

//func HandleRequest() {
//	r := mux.NewRouter()
//	r.Use(middleware.SetContentTypeMiddleware)
//	r.HandleFunc("/", controllers.Home).Methods("Get")
//	r.HandleFunc("/api/users", controllers.GetAllUsers).Methods("Get")
//	r.HandleFunc("/api/users/{username}/details", controllers.GetUserDetails).Methods("Get")
//	r.HandleFunc("/api/users/{username}/repos", controllers.GetUserRepositories).Methods("Get")
//
//	log.Fatal(http.ListenAndServe(":"+domain.GlobalConfig.Port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
//}

func (app *cmd.Application) Routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.SetAcceptMiddleware)

	// register routes
	mux.Get("/", controllers.Home)
	mux.Get("/api/users", controllers.GetAllUsers)
	mux.Get("/api/users/{username}/details", controllers.GetUserDetails)
	mux.Get("/api/users/{username}/repos", controllers.GetUserRepositories)

	// static assets
	fileServer := http.FileServer(http.Dir("./api/"))
	mux.Handle("/api/*", http.StripPrefix("/api", fileServer))
	return mux
}
