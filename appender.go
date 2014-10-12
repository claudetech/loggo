package loggo

import (
	"io"
	"os"
)

type appenderWithFilter struct {
	appender Appender
	filter   Filter
}

type Appender interface {
	Append(string)
}

type writerAppender struct {
	writer io.Writer
}

func (w *writerAppender) Append(s string) {
	_, _ = io.WriteString(w.writer, s)
}

func NewWriterAppender(writer io.Writer) Appender {
	return &writerAppender{writer: writer}
}

func NewStdoutAppender() Appender {
	return &writerAppender{writer: os.Stdout}
}
