package domain

// UseCaseFactory -
type UseCaseFactory struct {
	repository RepositoryGateway
}

// NewUseCaseFactory - constructor that injects the repository.
func NewUseCaseFactory(repository RepositoryGateway) *UseCaseFactory {
	return &UseCaseFactory{
		repository: repository,
	}
}
