package loggo

import (
	"strings"
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
	appenders  []*appenderWithFilter
	linebreak  string
	nowFunc    func() time.Time
	dateFormat string
	color      bool
	padding    bool
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

func (l *Logger) AddAppender(appender Appender) {
	l.addAppender(appender, nil, false)
}

func (l *Logger) AddColoredAppender(appender Appender) {
	l.addAppender(appender, nil, true)
}

func (l *Logger) AddAppenderWithFilter(appender Appender, filter Filter) {
	l.addAppender(appender, filter, false)
}

func (l *Logger) AddColoredAppenderWithFilter(appender Appender, filter Filter) {
	l.addAppender(appender, filter, true)
}

func (l *Logger) addAppender(appender Appender, filter Filter, color bool) {
	appenderContainer := &appenderWithFilter{
		appender: appender,
		filter:   filter,
		color:    color,
	}
	l.appenders = append(l.appenders, appenderContainer)
}

func (l *Logger) EnableColor() {
	l.color = true
}

func (l *Logger) DisableColor() {
	l.color = true
}

func (l *Logger) Verbose(content string) {
	l.Log(Verbose, content)
}

func (l *Logger) Debug(content string) {
	l.Log(Debug, content)
}

func (l *Logger) Info(content string) {
	l.Log(Info, content)
}

func (l *Logger) Warning(content string) {
	l.Log(Warning, content)
}

func (l *Logger) Error(content string) {
	l.Log(Error, content)
}

func (l *Logger) Critical(content string) {
	l.Log(Critical, content)
}

func (l *Logger) EnablePadding() {
	l.padding = true
}

func (l *Logger) DisablePadding() {
	l.padding = false
}

func (l *Logger) Log(level Level, content string) {
	msg := &Message{
		Name:       l.Name(),
		Level:      level,
		Content:    content,
		Time:       l.nowFunc(),
		dateFormat: l.DateFormat(),
		padding:    l.padding,
		tpl:        l.tpl,
	}
	l.log(msg)
}

func (l *Logger) log(msg *Message) {
	if msg.Level < l.Level() {
		return
	}

	for _, appenderContainer := range l.appenders {
		if appenderContainer.filter == nil || appenderContainer.filter.ShouldLog(msg) {
			msg.color = l.color && appenderContainer.color
			appenderContainer.appender.Append(msg)
		}
	}
}
