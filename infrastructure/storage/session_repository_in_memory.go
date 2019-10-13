package storage

import (
	"time"

	entity "github.com/beilmo/spectre-go-rest-api/domain/model"
)

type SessionRepositoryInMemory struct {
	List []entity.Session
}

// NewSessionRepositoryInMemory create new instance of SessionRepositoryInMemory
func NewSessionRepositoryInMemory() SessionRepositoryInMemory {
	var list []entity.Session
	list = append(list, entity.Session{
		ID:       1,
		Title:    "Hacking with Go",
		Abstract: "Such wow",
		Room:     "303",
		Date:     time.Now(),
		Duration: 45,
		Speakers: []entity.Speaker{},
		Keywords: []string{},
	})
	new := SessionRepositoryInMemory{}
	new.List = list
	return new
}

func (s SessionRepositoryInMemory) FindAll() ([]entity.Session, error) {
	return s.List, nil
}

func (s SessionRepositoryInMemory) FindByID(int64) (entity.Session, error) {
	return s.List[0], nil
}

func (s SessionRepositoryInMemory) Store(newSession entity.Session) (entity.Session, error) {
	s.List = append(s.List, newSession)
	return newSession, nil
}

func (SessionRepositoryInMemory) Update(newSession entity.Session) (entity.Session, error) {
	return newSession, nil
}

func (SessionRepositoryInMemory) Delete(item entity.Session) (entity.Session, error) {
	return item, nil
}
