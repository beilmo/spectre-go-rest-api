package logging

import (
	"log"
)

// ConsoleLogger type.
type ConsoleLogger struct{}

// Log - concrete implementation of function `Log` defined by `Log` interface.
func (logger ConsoleLogger) Log(args ...interface{}) {
	log.Println(args...)
}

// LogWarning -
func (logger ConsoleLogger) LogWarning(warning string) {
	log.Panic(warning)
}

// LogFatality -
func (logger ConsoleLogger) LogFatality(fatality error) {
	log.Fatalln(fatality)
}
