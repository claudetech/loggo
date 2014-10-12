package loggo_default

import (
	"github.com/claudetech/loggo"
)

var Log *loggo.Logger = func() *loggo.Logger {
	log := loggo.New("default")
	log.AddColoredAppenderWithFilter(loggo.NewStdoutAppender(), &loggo.MaxLogLevelFilter{MaxLevel: loggo.Info})
	log.AddColoredAppenderWithFilter(loggo.NewStderrAppender(), &loggo.MinLogLevelFilter{MinLevel: loggo.Warning})
	return log
}()
