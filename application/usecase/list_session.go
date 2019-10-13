package usecase

import (
	"github.com/beilmo/spectre-go-rest-api/application/repository"
	"github.com/beilmo/spectre-go-rest-api/domain/model"
)

// ListSessionOutputBoundary definition. To be implemented by presenters.
type ListSessionOutputBoundary interface {
	PresentSession(model.Session, error)
}

// ListSessionUseCase type. Implements Interactor as Input Boundary.
type ListSessionUseCase struct {
	SessionID  int64
	Repository repository.SessionRepository
	OutputPort ListSessionOutputBoundary
}

// Run - Interactor interface implementation
func (u ListSessionUseCase) Run(competion func()) {
	result, error := u.Repository.FindByID(u.SessionID)
	u.OutputPort.PresentSession(result, error)
	competion()
}
