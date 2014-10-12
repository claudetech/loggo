package loggo

import (
	"io"
	"os"
)

type appenderWithFilter struct {
	appender Appender
	filter   Filter
	color    bool
}

type Appender interface {
	Append(string, Level)
}

type writerAppender struct {
	writer io.Writer
}

func (w *writerAppender) Append(s string, level Level) {
	_, _ = io.WriteString(w.writer, s)
}

func NewWriterAppender(writer io.Writer) Appender {
	return &writerAppender{writer: writer}
}

func NewStdoutAppender() Appender {
	return &writerAppender{writer: os.Stdout}
}

func NewStderrAppender() Appender {
	return &writerAppender{writer: os.Stderr}
}
