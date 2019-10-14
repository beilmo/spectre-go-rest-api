package model

import "time"

// Session entity
type Session struct {
	ID       int64
	Title    string
	Abstract string
	Room     string
	Date     time.Time
	Duration int64
	Speakers []Speaker
	Keywords []string
}

func FilterSessions(list []Session, test func(Session) bool) (ret []Session) {
	for _, item := range list {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}

func SessionListContains(list []Session, test func(Session) bool) bool {
	for _, item := range list {
		if test(item) {
			return true
		}
	}
	return false
}
