package presentation

import (
	"github.com/beilmo/spectre-go-rest-api/domain"
)

// SessionViewModel type.
type SessionViewModel struct {
	useCaseFactory *domain.UseCaseFactory
}

// NewSessionViewModel - constructor that injects the useCaseFactory.
func NewSessionViewModel(useCaseFactory *domain.UseCaseFactory) *SessionViewModel {
	return &SessionViewModel{
		useCaseFactory: useCaseFactory,
	}
}
