package application

import (
	"github.com/beilmo/spectre-go-rest-api/application/repository"
)

type Storage struct {
	Session repository.SessionRepository
}
