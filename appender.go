package loggo

import (
	"io"
	"os"
	"sync"
)

const (
	// Color enabled flag
	Color = 1 << iota
	// Aync behavior flag
	Async = 1 << iota
)

type appenderContainer struct {
	appender Appender
	filter   Filter
	flags    int
	wlock    sync.Mutex
}

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

// Interface to append logs
type Appender interface {
	// Takes a message pointer and should
	// use it to log the message.
	Append(*Message)
}

type writerAppender struct {
	writer io.Writer
}

func (w *writerAppender) Append(msg *Message) {
	_, _ = io.WriteString(w.writer, msg.String())
}

// Creates a new appender that logs to the given io.Writer
func NewWriterAppender(writer io.Writer) Appender {
	return &writerAppender{writer: writer}
}

// Creates a new appender that logs to stdout
func NewStdoutAppender() Appender {
	return NewWriterAppender(stdout)
}

// Creates a new appender that logs to stderr
func NewStderrAppender() Appender {
	return NewWriterAppender(stderr)
}

// Creates a new appender that append logs to the given file
func NewFileAppender(path string) (Appender, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	return NewWriterAppender(f), err
}

func (w *writerAppender) Close() error {
	if closer, ok := w.writer.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
