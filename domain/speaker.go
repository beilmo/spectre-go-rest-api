package domain

// Speaker entity
type Speaker struct {
	ID          int64
	FirstName   string
	LastName    string
	JobTitle    string
	Company     string
	Biography   string
	PhotoURL    string
	FacebookURL string
	TwitterURL  string
	LinkedinURL string
	WebsiteURL  string
}

func FilterSpeakers(list []Speaker, test func(Speaker) bool) (ret []Speaker) {
	for _, item := range list {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}

func SpeakerListContains(list []Speaker, element Speaker) bool {
	for _, candidate := range list {
		if candidate == element {
			return true
		}
	}
	return false
}
