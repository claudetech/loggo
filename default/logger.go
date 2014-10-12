package loggo_default

import (
	"github.com/claudetech/loggo"
)

var Log *loggo.Logger = func() *loggo.Logger {
	log := loggo.New("default")
	log.AddAppender(loggo.NewStdoutAppender())
	return log
}()
