package usecase

import (
	"github.com/beilmo/spectre-go-rest-api/application"
	"github.com/beilmo/spectre-go-rest-api/domain/model"
)

// ListSessionOutputBoundary definition. To be implemented by presenters.
type ListSessionOutputBoundary interface {
	PresentSession(model.Session, error)
}

// ListSessionUseCase type. Implements Interactor as Input Boundary.
type ListSessionUseCase struct {
	SessionID  int64
	Repository application.Repository
	OutputPort ListSessionOutputBoundary
}

// Run - Interactor interface implementation
func (u ListSessionUseCase) Run(competion func()) {
	fail := func(e error) {
		u.OutputPort.PresentSession(model.Session{}, e)
		competion()
	}

	session, error := u.Repository.Session.FindByID(u.SessionID)
	if error != nil {
		fail(error)
		return
	}

	speakerIDs, error := u.Repository.SessionSpeakerLink.FindSpeakersBySession(u.SessionID)
	if error != nil {
		fail(error)
		return
	}

	speakers, error := u.Repository.Speaker.FindByIDs(speakerIDs)
	if error != nil {
		fail(error)
		return
	}

	// Put everything together.
	session.Speakers = speakers

	u.OutputPort.PresentSession(session, nil)
	competion()
}
