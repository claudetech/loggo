package loggo

import (
	"strings"
)

// Level represents representing the log level
type Level int32

// Constants representing each log level
const (
	Trace Level = iota
	Debug
	Info
	Warning
	Error
	Fatal
)

// Returns a string representation of the log level
func (l Level) String() string {
	switch l {
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Returns the log level from the passed string
// Returns Info if the string is not a correct log level
func LevelFromString(level string) Level {
	switch strings.ToLower(level) {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}
