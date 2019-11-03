package serializers

import (
	"time"

	"github.com/beilmo/spectre-go-rest-api/domain"
	"github.com/beilmo/spectre-go-rest-api/presentation/dto"
)

// DecodeSession - function that converts from a dto entity to a domain entity.
func DecodeSession(in dto.Session) domain.Session {
	out := domain.Session{
		ID:       in.Id,
		Title:    in.Title,
		Abstract: in.Abstract,
		Room:     in.Room,
		Date:     time.Unix(in.Date, 0),
		Duration: in.Duration,
		Keywords: in.Keywords,
	}

	return out
}

// EncodeSession - function that converts from a domain entity to a dto entity.
func EncodeSession(in domain.Session) dto.Session {
	speakers := []*dto.Speaker{}

	for _, speaker := range in.Speakers {
		encodedValue := EncodeSpeaker(speaker)
		speakers = append(speakers, &encodedValue)
	}

	out := dto.Session{
		Id:       in.ID,
		Title:    in.Title,
		Abstract: in.Abstract,
		Room:     in.Room,
		Date:     in.Date.Unix(),
		Duration: in.Duration,
		Speakers: speakers,
		Keywords: in.Keywords,
	}

	return out
}
