package main

import (
	"net/http"

	"github.com/beilmo/spectre-go-rest-api/application"

	"github.com/beilmo/spectre-go-rest-api/infrastructure/logging"
	"github.com/beilmo/spectre-go-rest-api/infrastructure/router"
	"github.com/beilmo/spectre-go-rest-api/infrastructure/storage/memory"
)

func main() {
	storage := application.Storage{
		Session: memory.NewSessionStorage(),
	}

	logger := logging.ConsoleLogger{}
	router := router.NewRouter(logger, storage)

	logger.Log("Serving requests on :8080")
	logger.LogFatality(http.ListenAndServe(":8088", router))
}
