package memory

import (
	"time"

	"github.com/beilmo/spectre-go-rest-api/domain/model"
	entity "github.com/beilmo/spectre-go-rest-api/domain/model"
)

// SessionStorage type.
type SessionStorage struct {
	List []entity.Session
}

// NewSessionStorage - creates new instance of SessionStorage.
func NewSessionStorage() SessionStorage {
	var list []entity.Session
	list = append(list, entity.Session{
		ID:       1,
		Title:    "Hacking with Go",
		Abstract: "Such wow",
		Room:     "303",
		Date:     time.Now(),
		Duration: 3600,
		Speakers: []entity.Speaker{},
		Keywords: []string{},
	})
	list = append(list, entity.Session{
		ID:       2,
		Title:    "Go by the stack overflow",
		Abstract: "Don't copy and paste",
		Room:     "301",
		Date:     time.Now(),
		Duration: 3000,
		Speakers: []entity.Speaker{entity.Speaker{ID: 12, FirstName: "Dorin", LastName: "Danciu"}},
		Keywords: []string{"tag1", "tag2", "foo", "bar", "golang"},
	})
	new := SessionStorage{}
	new.List = list
	return new
}

func (s SessionStorage) FindAll() ([]entity.Session, error) {
	return s.List, nil
}

func (s SessionStorage) FindByID(id int64) (entity.Session, error) {
	test := func(s entity.Session) bool { return s.ID == id }
	results := model.FilterSessions(s.List, test)

	if len(results) == 0 {
		return entity.Session{}, nil
	}

	return results[0], nil
}

func (s SessionStorage) Store(newSession entity.Session) (entity.Session, error) {
	s.List = append(s.List, newSession)
	return newSession, nil
}

func (SessionStorage) Update(newSession entity.Session) (entity.Session, error) {
	return newSession, nil
}

func (SessionStorage) Delete(item entity.Session) (entity.Session, error) {
	return item, nil
}
