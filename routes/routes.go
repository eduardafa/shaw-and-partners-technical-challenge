package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shaw-and-partners-technical-challenge/controllers"
	"shaw-and-partners-technical-challenge/domain"
	"shaw-and-partners-technical-challenge/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.SetContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/users", controllers.GetAllUsers).Methods("Get")
	r.HandleFunc("/api/users/{username}/details", controllers.GetUserDetails).Methods("Get")
	r.HandleFunc("/api/users/{username}/repos", controllers.GetUserRepositories).Methods("Get")

	log.Fatal(http.ListenAndServe(":"+domain.GlobalConfig.Port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
