package loggo_default

import (
	"github.com/claudetech/loggo"
	"os"
	"strings"
)

var Log *loggo.Logger = func() *loggo.Logger {
	name := strings.TrimLeft(strings.ToUpper(os.Args[0]), "./")
	log := loggo.New(name)
	log.AddColoredAppenderWithFilter(loggo.NewStdoutAppender(), &loggo.MaxLogLevelFilter{MaxLevel: loggo.Info})
	log.AddColoredAppenderWithFilter(loggo.NewStderrAppender(), &loggo.MinLogLevelFilter{MinLevel: loggo.Warning})
	return log
}()
