package main

import (
	"net/http"

	"github.com/beilmo/spectre-go-rest-api/domain"
	"github.com/beilmo/spectre-go-rest-api/persistence"
	"github.com/beilmo/spectre-go-rest-api/platform"
	"github.com/beilmo/spectre-go-rest-api/presentation"
)

func main() {
	/// 1. Create the REPO
	repo := persistence.NewInMemoryRepository()

	/// 2. Create the USE CASE FACTORY with the REPO
	useCaseFactory := domain.NewUseCaseFactory(repo)

	/// 3. Create the VIEW MODELS with the USE CASE FACTORY
	sessionViewModel := presentation.NewSessionViewModel(useCaseFactory)

	/// 4. Create the CONTROLLERS with the VIEW MODELS
	sessionController := platform.NewSessionController(sessionViewModel)

	/// 5. Create the ROUTER with the CONTROLLERS
	router := platform.NewRouter(sessionController)

	/// 6. Launch the SERVER
	http.ListenAndServe(":8080", router)
}
