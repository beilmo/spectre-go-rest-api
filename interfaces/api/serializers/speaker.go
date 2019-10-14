package serializers

import (
	"github.com/beilmo/spectre-go-rest-api/domain/model"
	"github.com/beilmo/spectre-go-rest-api/interfaces/api/dto"
)

// DecodeSpeaker - function that converts from a dto entity to a domain entity.
func DecodeSpeaker(in dto.Speaker) model.Speaker {
	out := model.Speaker{
		ID:          in.Id,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		JobTitle:    in.JobTitle,
		Company:     in.Company,
		Biography:   in.Biography,
		PhotoURL:    in.PhotoUrl,
		FacebookURL: in.FirstName,
		TwitterURL:  in.TwitterUrl,
		LinkedinURL: in.LinkedinUrl,
		WebsiteURL:  in.WebsiteUrl,
	}

	return out
}

// EncodeSpeaker - function that converts from a domain entity to a dto entity.
func EncodeSpeaker(in model.Speaker) dto.Speaker {
	out := dto.Speaker{
		Id:          in.ID,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		JobTitle:    in.JobTitle,
		Company:     in.Company,
		Biography:   in.Biography,
		PhotoUrl:    in.PhotoURL,
		FacebookUrl: in.FacebookURL,
		TwitterUrl:  in.TwitterURL,
		LinkedinUrl: in.LinkedinURL,
		WebsiteUrl:  in.WebsiteURL,
	}

	return out
}
