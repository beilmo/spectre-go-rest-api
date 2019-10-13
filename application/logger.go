package application

// Logger interface defining possible log functions.
type Logger interface {
	Log(args ...interface{})
	LogWarning(warning string)
	LogFatality(fatality error)
}
