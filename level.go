package loggo

// Type representing the log level
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
