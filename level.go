package loggo

type Level int32

const (
	Trace Level = iota
	Debug
	Info
	Warning
	Error
	Fatal
)

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
