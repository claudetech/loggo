package loggo

import (
	"strings"
	"text/template"
	"time"
)

const (
	defaultFormat     = "[{{.Name}}] [{{.TimeStr}}] {{.LevelStr}}: {{.Content}}\n"
	defaultDateFormat = "2006-02-01 15:04"
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
}

func New(name string) *Logger {
	logger := &Logger{
		level:      Debug,
		nowFunc:    time.Now,
		name:       name,
		linebreak:  "\n",
		dateFormat: defaultDateFormat,
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
	l.AddAppenderWithFilter(appender, nil)
}

func (l *Logger) AddAppenderWithFilter(appender Appender, filter Filter) {
	appenderContainer := &appenderWithFilter{
		appender: appender,
		filter:   filter,
	}
	l.appenders = append(l.appenders, appenderContainer)
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

func (l *Logger) Log(level Level, content string) {
	msg := &Message{
		Name:       l.Name(),
		Level:      level,
		Content:    content,
		Time:       l.nowFunc(),
		dateFormat: l.DateFormat(),
	}
	l.log(msg)
}

func (l *Logger) log(msg *Message) {
	if msg.Level < l.Level() {
		return
	}

	for _, appenderContainer := range l.appenders {
		if appenderContainer.filter == nil || appenderContainer.filter.ShouldLog(msg) {
			appenderContainer.appender.Append(msg.Format(l.tpl))
		}
	}
}
