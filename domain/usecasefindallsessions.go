package domain

// FindAllSessionsUseCase type.
type FindAllSessionsUseCase struct {
	Repository RepositoryGateway
}

// NewFindAllSessionsUseCase - constructor that injects the id, the repository and the output port.
func NewFindAllSessionsUseCase(r RepositoryGateway) *FindAllSessionsUseCase {
	// Return a new instance.
	return &FindAllSessionsUseCase{r}
}

// Execute - UseCase interface implementation
func (u *FindAllSessionsUseCase) Execute(ch chan<- Result) {
	fail := func(e error) {
		ch <- Result{[]Session{}, e}
	}

	succeed := func(s []Session) {
		ch <- Result{s, nil}
	}

	sessions, error := u.Repository.GetAllSessions()
	if error != nil {
		fail(error)
		return
	}

	succeed(sessions)
}
