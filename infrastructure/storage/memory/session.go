package memory

import (
	"time"

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
		ID:       1,
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

func (s SessionStorage) FindByID(int64) (entity.Session, error) {
	return s.List[0], nil
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
