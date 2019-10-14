package application

import (
	"github.com/beilmo/spectre-go-rest-api/application/repository"
)

// Repository type
type Repository struct {
	Session            repository.SessionRepository
	Speaker            repository.SpeakerRepository
	SessionSpeakerLink repository.SessionSpeakerLinkRepository
}
