package repository

import entity "github.com/beilmo/spectre-go-rest-api/domain/model"

// SessionSpeakerLinkRepository - Data access objects interface
type SessionSpeakerLinkRepository interface {
	FindSpeakersBySession(int64) ([]int64, error)
	FindSessionsBySpeaker(int64) ([]int64, error)
	LinkSpeakerToSession(speakerID int64, sessionID int64) (entity.SessionSpeakerLink, error)
	UnlinkSpeakerFromSession(speakerID int64, sessionID int64) (entity.SessionSpeakerLink, error)
}
