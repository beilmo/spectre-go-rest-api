package memory

import (
	entity "github.com/beilmo/spectre-go-rest-api/domain/model"
)

// SessionSpeakerLinkStorage type.
type SessionSpeakerLinkStorage struct {
	List []entity.SessionSpeakerLink
}

// NewSessionSpeakerLinkStorage - creates new instance of SessionSpeakerLinkStorage.
func NewSessionSpeakerLinkStorage() SessionSpeakerLinkStorage {
	var list []entity.SessionSpeakerLink
	list = append(list, entity.SessionSpeakerLink{
		SessionID: 1,
		SpeakerID: 1,
	})
	list = append(list, entity.SessionSpeakerLink{
		SessionID: 1,
		SpeakerID: 2,
	})
	list = append(list, entity.SessionSpeakerLink{
		SessionID: 2,
		SpeakerID: 2,
	})
	new := SessionSpeakerLinkStorage{}
	new.List = list
	return new
}

func (s SessionSpeakerLinkStorage) FindSpeakersBySession(id int64) ([]int64, error) {
	speakerIDs := []int64{}

	for _, link := range s.List {
		if link.SessionID == id {
			speakerIDs = append(speakerIDs, link.SpeakerID)
		}
	}

	return speakerIDs, nil
}

func (s SessionSpeakerLinkStorage) FindSessionsBySpeaker(int64) ([]int64, error) {
	return []int64{}, nil
}

func (s SessionSpeakerLinkStorage) LinkSpeakerToSession(speakerID int64, sessionID int64) (entity.SessionSpeakerLink, error) {
	return entity.SessionSpeakerLink{}, nil
}

func (s SessionSpeakerLinkStorage) UnlinkSpeakerFromSession(speakerID int64, sessionID int64) (entity.SessionSpeakerLink, error) {
	return entity.SessionSpeakerLink{}, nil
}
