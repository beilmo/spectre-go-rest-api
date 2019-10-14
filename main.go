package main

import (
	"net/http"

	"github.com/beilmo/spectre-go-rest-api/application"

	"github.com/beilmo/spectre-go-rest-api/infrastructure/logging"
	"github.com/beilmo/spectre-go-rest-api/infrastructure/router"
	"github.com/beilmo/spectre-go-rest-api/infrastructure/storage/memory"
)

func main() {
	repository := application.Repository{
		Session:            memory.NewSessionStorage(),
		Speaker:            memory.NewSpeakerStorage(),
		SessionSpeakerLink: memory.NewSessionSpeakerLinkStorage(),
	}

	logger := logging.ConsoleLogger{}
	router := router.NewRouter(logger, repository)

	logger.Log("Serving requests on :8080")
	logger.LogFatality(http.ListenAndServe(":8091", router))
}
