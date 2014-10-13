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

type Logger struct {
	name         string
	format       string
	tpl          *template.Template
	level        Level
	appenders    []*appenderContainer
	linebreak    string
	nowFunc      func() time.Time
	dateFormat   string
	color        bool
	padding      bool
	lockSettings bool
	callerInfo   bool
	wlock        sync.Mutex
}

func New(name string) *Logger {
	logger := &Logger{
		level:        Debug,
		nowFunc:      time.Now,
		name:         name,
		linebreak:    "\n",
		dateFormat:   defaultDateFormat,
		color:        true,
		padding:      true,
		callerInfo:   false,
		lockSettings: false,
	}
	logger.SetFormat(defaultFormat)
	loggers[name] = logger
	return logger
}

func (l *Logger) Name() string {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	return l.name
}

func (l *Logger) ThreasafeSettings() bool {
	return l.lockSettings
}

func (l *Logger) SetThreasafeSettings(b bool) {
	l.lockSettings = b
}

func (l *Logger) SetName(name string) {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.name = name
}

func (l *Logger) SetNowFunc(f func() time.Time) {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.nowFunc = f
}

func (l *Logger) Level() Level {
	return Level(atomic.LoadInt32((*int32)(&l.level)))
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func (l *Logger) Linebreak() string {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	return l.linebreak
}

func (l *Logger) SetLineBreak(linebreak string) {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.linebreak = linebreak
}

func (l *Logger) Format() string {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	return l.format
}

func (l *Logger) SetFormat(format string) error {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
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

func (l *Logger) DateFormat() string {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	return l.dateFormat
}

func (l *Logger) SetDateFormat(format string) {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.dateFormat = format
}

func (l *Logger) AddAppender(appender Appender, flags int) {
	l.AddAppenderWithFilter(appender, nil, flags)
}

func (l *Logger) AddAppenderWithFilter(appender Appender, filter Filter, flags int) {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	container := &appenderContainer{
		appender: appender,
		filter:   filter,
		flags:    flags,
	}
	l.appenders = append(l.appenders, container)
}

func (l *Logger) EnableColor() {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.color = true
}

func (l *Logger) Color() bool {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	return l.color
}

func (l *Logger) DisableColor() {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.color = true
}

func (l *Logger) EnablePadding() {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.padding = true
}

func (l *Logger) DisablePadding() {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}
	l.padding = false
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Logf(Trace, format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(Debug, format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(Info, format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Logf(Warning, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(Error, format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(Fatal, format, v...)
}

func (l *Logger) Trace(v ...interface{}) {
	l.Log(Trace, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Log(Debug, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Log(Info, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.Log(Warning, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Log(Error, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Log(Fatal, v...)
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

func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	msg := l.makeMessage(level, fmt.Sprintf(format, v...))
	l.log(msg)
}

func (l *Logger) Log(level Level, v ...interface{}) {
	msg := l.makeMessage(level, fmt.Sprint(v...))
	l.log(msg)
}

func (l *Logger) destroyAppender(appender Appender) error {
	if closer, ok := appender.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

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

func (l *Logger) makeAppend(container *appenderContainer, msg *Message) {
	if container.flags&NoLock == 0 {
		container.wlock.Lock()
		defer container.wlock.Unlock()
	}
	container.appender.Append(msg)
}

func (l *Logger) log(msg *Message) {
	if l.lockSettings {
		l.wlock.Lock()
		defer l.wlock.Unlock()
	}

	if msg.Level < l.level {
		return
	}

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
