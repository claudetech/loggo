package loggo

const EmptyFlag = 0

var Colors = map[Level]string{
	Trace:   "white",
	Debug:   "blue",
	Info:    "cyan",
	Warning: "yellow",
	Error:   "magenta",
	Fatal:   "red",
}

var loggers = make(map[string]*Logger)

func Get(name string) *Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}
	return nil
}
