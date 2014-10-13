package loggo

type Level int32

const (
	Verbose Level = iota
	Debug
	Info
	Warning
	Error
	Critical
)

func (l Level) String() string {
	switch l {
	case Verbose:
		return "VERBOSE"
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Critical:
		return "CRITIC"
	default:
		return "UNKNOWN"
	}
}
