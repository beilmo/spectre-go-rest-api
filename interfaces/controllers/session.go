package controllers

import (
	"github.com/beilmo/spectre-go-rest-api/application"
	"github.com/beilmo/spectre-go-rest-api/application/repository"
	"github.com/beilmo/spectre-go-rest-api/application/usecase"
)

// SessionController type.
type SessionController struct {
	Repository repository.SessionRepository
	Logger     application.Logger
}

// NewSessionController - factory function.
func NewSessionController(repository repository.SessionRepository, logger application.Logger) *SessionController {
	return &SessionController{
		Repository: repository,
		Logger:     logger,
	}
}

// Instance methods

// ListAllSessions - Sends all sessions through the output port.
func (c SessionController) ListAllSessions(o usecase.ListSessionsOutputBoundary) {
	uci := InteractorForListingAllSessions(c.Repository, o)
	uci.Run(func() {})
}

// ListSessionByID - Sends the required session through the output port.
func (c SessionController) ListSessionByID(id int64, o usecase.ListSessionOutputBoundary) {
	uci := InteractorForListingSessionsByID(id, c.Repository, o)
	uci.Run(func() {})
}

// Factory functions

// InteractorForListingAllSessions - Factory function creating the interactor responsible of resolving all sessions.
func InteractorForListingAllSessions(r repository.SessionRepository, o usecase.ListSessionsOutputBoundary) usecase.Interactor {
	return &usecase.ListSessionsUseCase{
		Repository: r,
		OutputPort: o,
	}
}

// InteractorForListingSessionsByID - Factory function creating the interactor responsible of resolving a particular session.
func InteractorForListingSessionsByID(id int64, r repository.SessionRepository, o usecase.ListSessionOutputBoundary) usecase.Interactor {
	return &usecase.ListSessionUseCase{
		SessionID:  id,
		Repository: r,
		OutputPort: o,
	}
}
