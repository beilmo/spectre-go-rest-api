package usecase

// Interactor interface to be impemented by all concrete usecases.
// Acts as Use Case Input Port interface.
type Interactor interface {
	Run(competion func())
}
