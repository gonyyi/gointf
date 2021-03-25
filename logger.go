package gointf

// Logger is an interface for standard level'd with additional tagging info as int64
type Logger interface {
	Trace(tag int64, action string, detail string)
	Debug(tag int64, action string, detail string)
	Info(tag int64, action string, detail string)
	Warn(tag int64, action string, detail string)
	Error(tag int64, action string, detail string)
	Fatal(tag int64, action string, detail string)
}
