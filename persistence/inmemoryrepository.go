package persistence

import (
	"time"

	"github.com/beilmo/spectre-go-rest-api/domain"
)

// InMemoryRepository -
type InMemoryRepository struct {
	sessions []domain.Session
}

// NewInMemoryRepository - constructor.
func NewInMemoryRepository() *InMemoryRepository {
	var sessions []domain.Session
	sessions = append(sessions, domain.Session{
		ID:       1,
		Title:    "Hacking with Go",
		Abstract: "Such wow",
		Room:     "303",
		Date:     time.Now(),
		Duration: 3600,
		Speakers: []domain.Speaker{},
		Keywords: []string{},
	})
	sessions = append(sessions, domain.Session{
		ID:       2,
		Title:    "Go by the stack overflow",
		Abstract: "Don't copy and paste",
		Room:     "301",
		Date:     time.Now(),
		Duration: 3000,
		Speakers: []domain.Speaker{domain.Speaker{ID: 12, FirstName: "Dorin", LastName: "Danciu"}},
		Keywords: []string{"tag1", "tag2", "foo", "bar", "golang"},
	})
	return &InMemoryRepository{sessions}
}

// GetAllSessions -
func (r InMemoryRepository) GetAllSessions() ([]domain.Session, error) {
	return r.sessions, nil
}

// FindSession -
func (r InMemoryRepository) FindSession(id int64) (domain.Session, error) {
	test := func(s domain.Session) bool { return s.ID == id }
	results := domain.FilterSessions(r.sessions, test)

	if len(results) == 0 {
		return domain.Session{}, nil
	}

	return results[0], nil
}
