package loggo

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"text/template"
	"time"
)

const (
	defaultFormat     = "[{{.Name}}] [{{.TimeStr}}] {{.LevelStr}}: {{.Content}}\n"
	defaultDateFormat = "2006-01-02 15:04"
)

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
	wlock      sync.Mutex
}

func New(name string) *Logger {
	logger := &Logger{
		level:      Debug,
		nowFunc:    time.Now,
		name:       name,
		linebreak:  "\n",
		dateFormat: defaultDateFormat,
		color:      true,
		padding:    true,
	}
	logger.SetFormat(defaultFormat)
	return logger
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) SetName(name string) {
	l.name = name
}

func (l *Logger) SetNowFunc(f func() time.Time) {
	l.nowFunc = f
}

func (l *Logger) Level() Level {
	return l.level
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) Linebreak() string {
	return l.linebreak
}

func (l *Logger) SetLineBreak(linebreak string) {
	l.linebreak = linebreak
}

func (b *Logger) Format() string {
	return b.format
}

func (l *Logger) SetFormat(format string) error {
	if !strings.HasSuffix(format, l.Linebreak()) {
		format += l.Linebreak()
	}
	tpl, err := template.New("loggerTemplate").Parse(format)
	if err != nil {
		return err
	}
	l.format = format
	l.tpl = tpl
	return nil
}

func (l *Logger) DateFormat() string {
	return l.dateFormat
}

func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

func (l *Logger) AddAppender(appender Appender, flags int) {
	l.AddAppenderWithFilter(appender, nil, flags)
}

func (l *Logger) AddAppenderWithFilter(appender Appender, filter Filter, flags int) {
	container := &appenderContainer{
		appender: appender,
		filter:   filter,
		flags:    flags,
	}
	l.appenders = append(l.appenders, container)
}

func (l *Logger) EnableColor() {
	l.color = true
}

func (l *Logger) DisableColor() {
	l.color = true
}

func (l *Logger) Verbosef(format string, v ...interface{}) {
	l.Logf(Verbose, format, v...)
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

func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.Logf(Critical, format, v...)
}

func (l *Logger) Verbose(v ...interface{}) {
	l.Log(Verbose, v...)
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

func (l *Logger) Critical(v ...interface{}) {
	l.Log(Critical, v...)
}

func (l *Logger) EnablePadding() {
	l.padding = true
}

func (l *Logger) DisablePadding() {
	l.padding = false
}

func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.Log(level, msg)
}

func (l *Logger) Log(level Level, v ...interface{}) {
	msg := &Message{
		Name:       l.Name(),
		Level:      level,
		Content:    fmt.Sprint(v...),
		Time:       l.nowFunc(),
		dateFormat: l.DateFormat(),
		padding:    l.padding,
		tpl:        l.tpl,
	}
	l.log(msg)
}

func (l *Logger) destroyAppender(appender Appender) error {
	if closer, ok := appender.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (l *Logger) Destroy() (err error) {
	for _, container := range l.appenders {
		if e := l.destroyAppender(container.appender); e != nil {
			err = e
		}
	}
	l.appenders = nil
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
	if msg.Level < l.Level() {
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
