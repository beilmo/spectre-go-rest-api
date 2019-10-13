package usecase

import (
	"github.com/beilmo/spectre-go-rest-api/application/repository"
	"github.com/beilmo/spectre-go-rest-api/domain/model"
)

// ListSessionsOutputBoundary definition. To be implemented by presenters.
type ListSessionsOutputBoundary interface {
	PresentSessionsList([]model.Session, error)
}

// ListSessionsUseCase type. Implements Interactor as Input Boundary.
type ListSessionsUseCase struct {
	Repository repository.SessionRepository
	OutputPort ListSessionsOutputBoundary
}

// Run - Interactor interface implementation
func (u ListSessionsUseCase) Run(competion func()) {
	result, error := u.Repository.FindAll()
	u.OutputPort.PresentSessionsList(result, error)
	competion()
}
