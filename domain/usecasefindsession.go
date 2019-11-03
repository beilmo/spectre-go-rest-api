package domain

// FindSessionUseCase type.
type FindSessionUseCase struct {
	SessionID  int64
	Repository RepositoryGateway
}

// NewFindSessionUseCase - constructor that injects the id, the repository and the output port.
func NewFindSessionUseCase(id int64, r RepositoryGateway) *FindSessionUseCase {
	// Return a new instance.
	return &FindSessionUseCase{id, r}
}

// Execute - UseCase interface implementation
func (u *FindSessionUseCase) Execute(ch chan<- Result) {
	fail := func(e error) {
		ch <- Result{Session{}, e}
	}

	succeed := func(s Session) {
		ch <- Result{s, nil}
	}

	session, error := u.Repository.FindSession(u.SessionID)
	if error != nil {
		fail(error)
		return
	}

	succeed(session)
}
