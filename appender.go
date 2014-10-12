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
	Append(*Message)
}

type writerAppender struct {
	writer io.Writer
}

func (w *writerAppender) Append(msg *Message) {
	_, _ = io.WriteString(w.writer, msg.String())
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
