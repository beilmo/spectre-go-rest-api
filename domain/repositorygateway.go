package domain

type RepositoryGateway interface {
	GetAllSessions() ([]Session, error)
	FindSession(int64) (Session, error)
	// StoreSession(Session) (Session, error)
	// UpdateSession(Session) (Session, error)
	// DeleteSession(Session) (Session, error)

	// SpeakersLinkedToSession(int64) ([]Speaker, error)
	// LinkSpeakerToSession(speakerID int64, sessionID int64) (Session, error)
	// UnlinkSpeakerFromSession(speakerID int64, sessionID int64) (Session, error)
}
