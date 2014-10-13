// Package loggo provides easy to use
// and extensible logging functionality
package loggo

// Flag representing all options turned off
const EmptyFlag = 0

// Mapping between log level and color
// names used to display ANSI colors in terminal
// See https://github.com/mgutz/ansi for more info about accepted values
var Colors = map[Level]string{
	Trace:   "white",
	Debug:   "blue",
	Info:    "cyan",
	Warning: "yellow",
	Error:   "magenta",
	Fatal:   "red",
}

var loggers = make(map[string]*Logger)

// Retreive the logger with the given name
// Return nil if no such logger exists
func Get(name string) *Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}
	return nil
}
