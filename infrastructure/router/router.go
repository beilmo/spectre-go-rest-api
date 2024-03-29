package router

import (
	"fmt"
	"net/http"

	"github.com/beilmo/spectre-go-rest-api/application"
	"github.com/gorilla/mux"
)

// NewRouter fctory.
func NewRouter(logger application.Logger, repository application.Repository) *mux.Router {
	sessionsHandler := SessionsRequestHandler{
		Logger:     logger,
		Repository: repository,
	}

	router := mux.NewRouter().StrictSlash(true)
	// router.Use(security.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/sessions", sessionsHandler.FindAllSessions).Methods("GET")
	router.HandleFunc("/api/sessions/{id:[0-9]+}", sessionsHandler.FindSessionByID).Methods("GET")

	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
