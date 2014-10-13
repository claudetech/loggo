package loggo

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"
)

const (
	defaultFormat     = "[{{.NameUp}}] [{{.TimeStr}}] {{.LevelStr}}: {{.Content}}\n"
	defaultDateFormat = "2006-01-02 15:04"
)

// Logger is the basic struct for all logging operations
type Logger struct {
	name       string
	format     string
	tpl        *template.Template
	level      Level
	appenders  []*appenderContainer
	linebreak  string
	nowFunc    func() time.Time
	dateFormat string
	color      bool
	padding    bool
	callerInfo bool
	wlock      sync.Mutex
}

// New creates a new logger and registers it.
// The logger can then either be used directly
// or retreived using the name passed as argument.
func New(name string) *Logger {
	logger := &Logger{
		level:      Debug,
		nowFunc:    time.Now,
		name:       name,
		linebreak:  "\n",
		dateFormat: defaultDateFormat,
		color:      true,
		padding:    true,
		callerInfo: false,
	}
	logger.SetFormat(defaultFormat)
	loggers[name] = logger
	return logger
}

// Name returns the name of the logger
func (l *Logger) Name() string {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	return l.name
}

// SetName set the name of the logger
func (l *Logger) SetName(name string) {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.name = name
}

// SetNowFunc set the function used to get the current time.
// Defaults to time.Now
func (l *Logger) SetNowFunc(f func() time.Time) {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.nowFunc = f
}

// Level returns the current log level
func (l *Logger) Level() Level {
	return Level(atomic.LoadInt32((*int32)(&l.level)))
}

// SetLevel set the current log level
func (l *Logger) SetLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

// Linebreak returns the string used as linebreak
func (l *Logger) Linebreak() string {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	return l.linebreak
}

// SetLineBreak set the string to use as linebreak
// Defaults to "\n"
func (l *Logger) SetLineBreak(linebreak string) error {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	strings.TrimRight(l.format, l.linebreak)
	l.format += linebreak
	l.linebreak = linebreak
	return l.updateTemplate(l.format)
}

func (l *Logger) updateTemplate(format string) error {
	if !strings.HasSuffix(format, l.linebreak) {
		format += l.linebreak
	}
	tpl, err := template.New("loggerTemplate").Parse(format)
	if err != nil {
		return err
	}
	l.format = format
	l.tpl = tpl
	l.callerInfo = false
	for _, str := range []string{"{{.Line}}", "{{.File}}", "{{.FuncName}}"} {
		if strings.Contains(l.format, str) {
			l.callerInfo = true
			break
		}
	}
	return nil
}

// Format returns the current format
func (l *Logger) Format() string {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	return l.format
}

// SetFormat set the current format
// Defaults to "[{{.NameUp}}] [{{.TimeStr}}] {{.LevelStr}}: {{.Content}}"
func (l *Logger) SetFormat(format string) error {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	return l.updateTemplate(format)
}

// DateFormat returns the date format
func (l *Logger) DateFormat() string {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	return l.dateFormat
}

// SetDateFormat set the date format
// Defaults to "2006-01-02 15:04"
func (l *Logger) SetDateFormat(format string) {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.dateFormat = format
}

// AddAppender adds an appender to the logger
func (l *Logger) AddAppender(appender Appender, flags int) {
	l.AddAppenderWithFilter(appender, nil, flags)
}

// AddAppenderWithFilter adds an appender with a filter to the logger
func (l *Logger) AddAppenderWithFilter(appender Appender, filter Filter, flags int) {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	container := &appenderContainer{
		appender: appender,
		filter:   filter,
		flags:    flags,
	}
	l.appenders = append(l.appenders, container)
}

// Color returns the current status for global color
func (l *Logger) Color() bool {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	return l.color
}

// EnableColor enables color globally.
// Will allow appenders added with the `color` option
// to use colors.
func (l *Logger) EnableColor() {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.color = true
}

// DisableColor disables color globally.
// Event appenders added with the `color` option
// will not use colors.
func (l *Logger) DisableColor() {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.color = true
}

// EnablePadding enables padding so that all log level
// strings print with the same length.
func (l *Logger) EnablePadding() {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.padding = true
}

// DisablePadding disables padding
func (l *Logger) DisablePadding() {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	l.padding = false
}

// Tracef formats the given interfaces and logs with Trace level
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.logf(Trace, format, v...)
}

// Debugf formats the given interfaces and logs with Debug level
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logf(Debug, format, v...)
}

// Infof formats the given interfaces and logs with Info level
func (l *Logger) Infof(format string, v ...interface{}) {
	l.logf(Info, format, v...)
}

// Warningf formats the given interfaces and logs with Warning level
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.logf(Warning, format, v...)
}

// Errorf formats the given interfaces and logs with Error level
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logf(Error, format, v...)
}

// Fatalf formats the given interfaces and logs with Fatal level
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logf(Fatal, format, v...)
}

// Trace fogs the the given interfaces with Trace level
func (l *Logger) Trace(v ...interface{}) {
	l.log(Trace, v...)
}

// Debug fogs the the given interfaces with Debug level
func (l *Logger) Debug(v ...interface{}) {
	l.log(Debug, v...)
}

// Info fogs the the given interfaces with Info level
func (l *Logger) Info(v ...interface{}) {
	l.log(Info, v...)
}

// Warning fogs the the given interfaces with Warning level
func (l *Logger) Warning(v ...interface{}) {
	l.log(Warning, v...)
}

// Error fogs the the given interfaces with Error level
func (l *Logger) Error(v ...interface{}) {
	l.log(Error, v...)
}

// Fatal fogs the the given interfaces with Fatal level
func (l *Logger) Fatal(v ...interface{}) {
	l.log(Fatal, v...)
}

func (l *Logger) makeMessage(level Level, str string) *Message {
	msg := &Message{
		Name:       l.Name(),
		Level:      level,
		Content:    str,
		Time:       l.nowFunc(),
		dateFormat: l.DateFormat(),
		padding:    l.padding,
		tpl:        l.tpl,
	}
	if l.callerInfo {
		if pc, file, line, ok := runtime.Caller(3); ok {
			msg.File = file
			msg.Line = line
			if f := runtime.FuncForPC(pc); f != nil {
				msg.FuncName = f.Name()
			}
		}
	}
	return msg
}

// Logf formats interfaces with the given format and logs them with the given level
func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	l.logf(level, format, v...)
}

// Log logs the interfaces with the given level
func (l *Logger) Log(level Level, v ...interface{}) {
	l.log(level, v...)
}

func (l *Logger) logf(level Level, format string, v ...interface{}) {
	if level < l.Level() {
		return
	}
	msg := l.makeMessage(level, fmt.Sprintf(format, v...))
	l.outputLog(msg)
}

func (l *Logger) log(level Level, v ...interface{}) {
	if level < l.Level() {
		return
	}
	msg := l.makeMessage(level, fmt.Sprint(v...))
	l.outputLog(msg)
}

func (l *Logger) outputLog(msg *Message) {
	l.wlock.Lock()
	defer l.wlock.Unlock()

	for _, container := range l.appenders {
		if container.filter == nil || container.filter.ShouldLog(msg) {
			msg.color = l.color && (container.flags&Color != 0)
			if container.flags&Async == 0 {
				l.makeAppend(container, msg)
			} else {
				go l.makeAppend(container, msg)
			}
		}
	}
}

func (l *Logger) makeAppend(container *appenderContainer, msg *Message) {
	container.wlock.Lock()
	defer container.wlock.Unlock()
	container.appender.Append(msg)
}

func (l *Logger) destroyAppender(appender Appender) error {
	if closer, ok := appender.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// Destroy destroy the loggers, closing every appender implementing
// the io.Closer interface
func (l *Logger) Destroy() (err error) {
	l.wlock.Lock()
	defer l.wlock.Unlock()
	for _, container := range l.appenders {
		if e := l.destroyAppender(container.appender); e != nil {
			err = e
		}
	}
	l.appenders = nil
	delete(loggers, l.name)
	return
}
