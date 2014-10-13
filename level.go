package loggo

// Type representing the log level
type Level int32

const (
	// Trace log level
	Trace Level = iota
	// Debug log level
	Debug
	// Info log level
	Info
	// Warning log level
	Warning
	// Error log level
	Error
	// Fatal log level
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
