package domain

// Result type
type Result struct {
	Value interface{}
	Error error
}

// UseCase interface to be impemented by all concrete usecases.
type UseCase interface {
	Execute(ch chan<- Result)
}

// UseCaseFactory -
type UseCaseFactory struct {
	Repository RepositoryGateway
}

// NewUseCaseFactory - constructor that injects the repository.
func NewUseCaseFactory(r RepositoryGateway) *UseCaseFactory {
	return &UseCaseFactory{r}
}

// MakeAllSessions - factory method creating the UseCase responsible to provide
// all the sessions.
func (ucf *UseCaseFactory) MakeAllSessions() UseCase {
	return NewFindAllSessionsUseCase(ucf.Repository)
}

// MakeSessionByID - factory method creating the UseCase responsible to provide
// complete session details for a given id.
func (ucf *UseCaseFactory) MakeSessionByID(id int64) UseCase {
	return NewFindSessionUseCase(id, ucf.Repository)
}
