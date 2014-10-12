package loggo

import (
	"io"
	"os"
	"sync"
)

const (
	Color  = 1 << iota
	Async  = 1 << iota
	NoLock = 1 << iota
)

type appenderContainer struct {
	appender Appender
	filter   Filter
	flags    int
	wlock    sync.Mutex
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
