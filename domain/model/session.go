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
