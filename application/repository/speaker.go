package repository

import entity "github.com/beilmo/spectre-go-rest-api/domain/model"

// SpeakerRepository - Data access objects interface
type SpeakerRepository interface {
	FindAll() ([]entity.Speaker, error)
	FindByID(int64) (entity.Speaker, error)
	FindByIDs([]int64) ([]entity.Speaker, error)
	Store(entity.Speaker) (entity.Speaker, error)
	Update(entity.Speaker) (entity.Speaker, error)
	Delete(entity.Speaker) (entity.Speaker, error)
}
