package repository

import entity "github.com/beilmo/spectre-go-rest-api/domain/model"

// SessionRepository - Data access objects interface
type SessionRepository interface {
	FindAll() ([]entity.Session, error)
	FindByID(int64) (entity.Session, error)
	Store(entity.Session) (entity.Session, error)
	Update(entity.Session) (entity.Session, error)
	Delete(entity.Session) (entity.Session, error)
}
