package memory

import (
	entity "github.com/beilmo/spectre-go-rest-api/domain/model"
)

// SpeakerStorage type.
type SpeakerStorage struct {
	List []entity.Speaker
}

// NewSpeakerStorage - creates new instance of SpeakerStorage.
func NewSpeakerStorage() SpeakerStorage {
	var list []entity.Speaker
	list = append(list, entity.Speaker{
		ID:        1,
		FirstName: "Jonny",
		LastName:  "Cage",
		JobTitle:  "Soldier",
		Company:   "Army",
		Biography: "Mortal Kombat 11",
		PhotoURL:  "https://vignette.wikia.nocookie.net/muc/images/3/36/Kage.jpg/revision/latest/scale-to-width-down/1000?cb=20190215162133",
	})
	list = append(list, entity.Speaker{
		ID:        2,
		FirstName: "Shao",
		LastName:  "Kahn",
		JobTitle:  "Soldier",
		Company:   "Army",
		Biography: "Shao Kahn is a video game character introduced in Mortal Kombat II, and is a recurring character and antagonist of the video game series and extended franchise. ",
		PhotoURL:  "https://static.tvtropes.org/pmwiki/pub/images/shao_kahn_5.png",
	})
	new := SpeakerStorage{}
	new.List = list
	return new
}

func (s SpeakerStorage) FindAll() ([]entity.Speaker, error) {
	return s.List, nil
}

func (s SpeakerStorage) FindByID(id int64) (entity.Speaker, error) {
	test := func(s entity.Speaker) bool { return s.ID == id }
	results := entity.FilterSpeakers(s.List, test)

	if len(results) == 0 {
		return entity.Speaker{}, nil
	}

	return results[0], nil
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s SpeakerStorage) FindByIDs(ids []int64) ([]entity.Speaker, error) {
	test := func(s entity.Speaker) bool { return contains(ids, s.ID) }
	results := entity.FilterSpeakers(s.List, test)

	if len(results) == 0 {
		return []entity.Speaker{}, nil
	}

	return results, nil
}

func (s SpeakerStorage) Store(newSpeaker entity.Speaker) (entity.Speaker, error) {
	s.List = append(s.List, newSpeaker)
	return newSpeaker, nil
}

func (SpeakerStorage) Update(newSpeaker entity.Speaker) (entity.Speaker, error) {
	return newSpeaker, nil
}

func (SpeakerStorage) Delete(item entity.Speaker) (entity.Speaker, error) {
	return item, nil
}
