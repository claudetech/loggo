package loggo

type Level int

const (
	Verbose Level = iota
	Debug
	Info
	Warning
	Error
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
	default:
		return "UNKNOW"
	}
}
