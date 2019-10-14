package usecase

import (
	"github.com/beilmo/spectre-go-rest-api/application"
	"github.com/beilmo/spectre-go-rest-api/domain/model"
)

// ListSessionsOutputBoundary definition. To be implemented by presenters.
type ListSessionsOutputBoundary interface {
	PresentSessionsList([]model.Session, error)
}

// ListSessionsUseCase type. Implements Interactor as Input Boundary.
type ListSessionsUseCase struct {
	Repository application.Repository
	OutputPort ListSessionsOutputBoundary
}

// Run - Interactor interface implementation
func (u ListSessionsUseCase) Run(competion func()) {
	fail := func(e error) {
		u.OutputPort.PresentSessionsList([]model.Session{}, e)
		competion()
	}

	sessions, error := u.Repository.Session.FindAll()
	if error != nil {
		fail(error)
		return
	}

	n := len(sessions)
	for i := 0; i < n; i++ {
		sessionID := sessions[i].ID
		speakerIDs, error := u.Repository.SessionSpeakerLink.FindSpeakersBySession(sessionID)
		if error != nil {
			fail(error)
			return
		}

		speakers, error := u.Repository.Speaker.FindByIDs(speakerIDs)
		if error != nil {
			fail(error)
			return
		}

		sessions[i].Speakers = speakers
	}

	u.OutputPort.PresentSessionsList(sessions, nil)
	competion()
}
