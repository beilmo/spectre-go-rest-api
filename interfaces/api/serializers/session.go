package serializers

import (
	"time"

	"github.com/beilmo/spectre-go-rest-api/domain/model"
	"github.com/beilmo/spectre-go-rest-api/interfaces/api/dto"
)

// DecodeSession - function that converts from a dto entity to a domain entity.
func DecodeSession(in dto.Session) model.Session {
	out := model.Session{
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
func EncodeSession(in model.Session) dto.Session {
	out := dto.Session{
		Id:       in.ID,
		Title:    in.Title,
		Abstract: in.Abstract,
		Room:     in.Room,
		Date:     in.Date.Unix(),
		Duration: in.Duration,
		Keywords: in.Keywords,
	}

	return out
}
